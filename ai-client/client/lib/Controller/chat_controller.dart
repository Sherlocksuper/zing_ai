import 'dart:async';
import 'dart:convert';
import 'package:client/Controller/web_socket.dart';
import 'package:client/model/chat_struct.dart';
import 'package:client/model/message.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import '../Constant.dart';
import '../config.dart';
import 'user_controller.dart';

//因为涉及到刷新页面，所以使用GetsController
class ChatController extends GetxController {
  bool isEdited = false;
  List<ChatDetailStruct> chatList = [];
  ScrollController scrollController = ScrollController();

  //配置AI Bot
  void configAI() {
    String title = "默认标题";
    String systemMessage = "你是我的AI助手，我需要你的帮助";
    AIType aiType = AIType.Text;
    Get.defaultDialog(
      title: 'Config your AI',
      titleStyle: TextStyle(
        fontSize: 18,
        fontWeight: FontWeight.bold,
        color: Colors.grey[300],
      ),
      middleTextStyle: TextStyle(
        fontSize: 16,
        color: Colors.grey[100],
      ),
      backgroundColor: Colors.grey[850],
      radius: 8.0,
      barrierDismissible: true,
      content: SizedBox(
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Icon(
                  Icons.android,
                  color: Colors.cyanAccent,
                  size: 60,
                ),
                TextField(
                  decoration: InputDecoration(
                    hintText: title,
                    hintStyle: TextStyle(
                      color: Colors.grey[300],
                    ),
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  onChanged: (value) {
                    title = value;
                  },
                  style: TextStyle(
                    color: Colors.grey[300],
                  ),
                ),
                const SizedBox(height: 16),
                GetBuilder<ChatController>(
                  id: 'aiConfig',
                  builder: (_) {
                    return Row(
                      children: [
                        Text(
                          'AI Type:',
                          style: TextStyle(
                            color: Colors.grey[300],
                          ),
                        ),
                        const SizedBox(width: 16),
                        Radio(
                          value: AIType.Text,
                          groupValue: aiType,
                          onChanged: (value) {
                            aiType = value as AIType;
                            _.update(['aiConfig']);
                          },
                        ),
                        Text(
                          'Text',
                          style: TextStyle(
                            color: Colors.grey[300],
                          ),
                        ),
                        const SizedBox(width: 16),
                        Radio(
                          value: AIType.Image,
                          groupValue: aiType,
                          onChanged: (value) {
                            aiType = value as AIType;
                            _.update(['aiConfig']);
                          },
                        ),
                        Text(
                          'Image',
                          style: TextStyle(
                            color: Colors.grey[300],
                          ),
                        ),
                      ],
                    );
                  },
                ),
                Text(
                  'Would you like to start a chat now?',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontSize: 14,
                    color: Colors.grey[100],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
      textConfirm: 'YES',
      textCancel: 'NO',
      confirmTextColor: Colors.white,
      buttonColor: Colors.cyanAccent,
      cancelTextColor: Colors.grey[300],
      onConfirm: () async {
        Get.back();
        await EasyLoading.show(status: 'is adding...');
        await startAChat(title, systemMessage, aiType);
        await EasyLoading.dismiss();
      },
      onCancel: () {
        Get.back();
      },
    );
  }

  ///开始一个chat
  Future<void> startAChat(String title, String systemMessage, AIType type) async {
    String robotType = type.toString().split(".")[1];
    var map = {'title': title, /*'systemMessage': systemMessage, */ 'userId': UserController.me.id, "type": robotType};
    var response = await dio.post(Constant.StartAChatHAT, data: map);
    if (response.data["code"] == 200) {
      getChatList();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
    update();
    print(response);
  }

  ///发送消息
  Future<void> sendMessage(int chatId, String content) async {
    WSController.judgeNull();

    addMessage(chatId, "user", content);
    update([chatId]);

    dio.post(
      Constant.SENDMESSAGE,
      data: {'chatId': chatId, 'content': content},
      onReceiveProgress: (int count, int total) {
        print("count:$count,total:$total");
      },
      options: Options(
        responseType: ResponseType.stream,
      ),
    ).then(
      (value) {
        value.data.stream.listen((data) {
          String answer = Utf8Decoder().convert(data);
          receiveStreamingMessage(chatId, answer);
        });
      },
    );
  }

  Future<void> sendForImage(int chatId, String content) async {
    WSController.judgeNull();

    addMessage(chatId, "user", content);
    update([chatId]);

    var response = await dio.post(
      Constant.SENDFORIMAGE,
      data: {'chatId': chatId, 'content': content},
    );

    print(response);

    if (response.data["code"] == 200) {
      addMessage(chatId, "assistant", response.data["data"]);
    } else {
      EasyLoading.showError(response.data["message"]);
    }

    update([chatId]);
  }

  ///获取消息列表
  Future<void> getChatList() async {
    try {
      var response = await dio.get("${Constant.GETCHATLIST}?userId=${UserController.me.id}");
      chatList = (response.data["code"] == 200 || response.data["data"] != null)
          ? List<ChatDetailStruct>.from(response.data["data"].map((e) => ChatDetailStruct.fromJson(e)))
          : throw Exception(response.data["message"]);
      print(chatList.length);
    } catch (e) {
      print(e);
    }
    update();
  }

  ///删除chat
  Future<void> deleteChat(int chatId) async {
    var response = await dio.get("${Constant.DELETECHAT}?chatId=$chatId");
    if (response.data["code"] == 200) {
      chatList.removeWhere((element) => element.id == chatId);
      update();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
  }

  ///清空chat
  Future<void> clearChat() async {
    var response = await dio.get("${Constant.DELETEALLCHAT}?userId=${UserController.me.id}");
    print(response.data);
    if (response.data["code"] == 200) {
      chatList.clear();
      update();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
  }

  void addMessage(int chatId, String role, String message) {
    print("正在添加数据");
    chatList
        .firstWhere((element) => element.id == chatId)
        .messages
        .insert(0, Message(chatId: chatId, role: role, content: message, createdAt: DateTime.now().toString()));

    print("添加$role的信息完成，正在滚动");

    update([chatId]);

    Timer(const Duration(milliseconds: 100), () {
      scrollController.position.animateTo(
        scrollController.position.minScrollExtent,
        duration: const Duration(milliseconds: 300),
        curve: Curves.easeOut,
      );
    });
  }

  void receiveStreamingMessage(int chatId, String message) {
    if (chatList.firstWhere((element) => element.id == chatId).messages.first.role == "user") {
      addMessage(chatId, "assistant", message);
    } else {
      chatList.firstWhere((element) => element.id == chatId).messages.first.content += message;
    }

    update([chatId]);
  }
}
