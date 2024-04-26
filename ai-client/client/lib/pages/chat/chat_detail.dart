import 'dart:async';

import 'package:client/model/chat_struct.dart';
import 'package:client/model/message.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';
import '../../Controller/chat_controller.dart';

class ChatDetail extends StatelessWidget {
  final ChatDetailStruct chat;
  final textInputController = TextEditingController();

  ChatDetail({super.key, required this.chat});

  @override
  Widget build(BuildContext context) {
    return GetBuilder<ChatController>(
      id: chat.id,
      builder: (logic) {
        Timer(const Duration(milliseconds: 100), () {
          logic.scrollController.jumpTo(
            logic.scrollController.position.minScrollExtent,
          );
        });
        return Scaffold(
          appBar: AppBar(
            title: Text(chat.title),
          ),
          body: Column(
            children: [
              Expanded(
                child: Stack(
                  children: [
                    ListView.builder(
                      shrinkWrap: true,
                      reverse: true,
                      physics: const ClampingScrollPhysics(),
                      controller: logic.scrollController,
                      itemBuilder: (context, index) {
                        return ChatBubbles(
                          messageStruct: chat.messages[index],
                          type: chat.type,
                        );
                      },
                      itemCount: chat.messages.length,
                    ),
                    Positioned(
                      top: 10,
                      right: 10,
                      child: FloatingActionButton(
                        onPressed: () {
                          logic.scrollController.animateTo(
                            logic.scrollController.position.minScrollExtent,
                            duration: const Duration(milliseconds: 500),
                            curve: Curves.easeInOut,
                          );
                        },
                        mini: true,
                        child: const Icon(Icons.arrow_downward),
                      ),
                    ),
                  ],
                ),
              ),
              Container(
                height: 50,
                padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                decoration: BoxDecoration(
                  color: Colors.grey[200], // Light grey color for the background
                  boxShadow: [
                    BoxShadow(
                      color: Colors.grey.withOpacity(0.5), // Shadow color with some transparency
                      spreadRadius: 2,
                      blurRadius: 4,
                      offset: const Offset(0, 3), // Position of the shadow
                    ),
                  ],
                ),
                child: Row(
                  children: <Widget>[
                    Expanded(
                      child: TextField(
                        controller: textInputController,
                        decoration: InputDecoration(
                          contentPadding: const EdgeInsets.symmetric(vertical: 5.0, horizontal: 10.0),
                          hintText: '输入消息...',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(20.0), // Rounded corners for the input field
                          ),
                          fillColor: Colors.white,
                          filled: true,
                        ),
                        onChanged: (value) {
                          textInputController.text = value;
                        },
                      ),
                    ),
                    IconButton(
                      icon: const Icon(Icons.send, color: Colors.blue), // Icon color changed to blue
                      onPressed: () async {
                        var content = textInputController.text;
                        textInputController.clear();

                        FocusScope.of(context).requestFocus(FocusNode());

                        chat.type == 'Image'
                            ? await logic.sendForImage(chat.id, content)
                            : await logic.sendMessage(chat.id, content);

                        logic.scrollController.animateTo(
                          logic.scrollController.position.maxScrollExtent,
                          duration: const Duration(milliseconds: 500),
                          curve: Curves.easeInOut,
                        );
                      },
                    ),
                  ],
                ),
              )
            ],
          ),
        );
      },
    );
  }
}

class ChatBubbles extends StatelessWidget {
  final String type;
  final Message messageStruct;

  const ChatBubbles({super.key, required this.messageStruct, required this.type});

  @override
  Widget build(BuildContext context) {
    final message = messageStruct.content;
    final role = messageStruct.role;
    return Align(
      alignment: role == "user"
          ? Alignment.centerRight
          : role == "system"
              ? Alignment.center
              : Alignment.centerLeft,
      child: Container(
        constraints: BoxConstraints(maxWidth: MediaQuery.of(context).size.width * (2 / 3)),
        padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 10),
        margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
        decoration: BoxDecoration(
          gradient: _getBubbleGradient(role),
          borderRadius: _getBubbleBorderRadius(role),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(0.2),
              spreadRadius: 1,
              blurRadius: 3,
              offset: const Offset(0, 1),
            ),
          ],
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            type == 'Text' || messageStruct.role == 'user'
                ? MarkdownBody(
                    data: message,
                    selectable: true,
                    styleSheet: MarkdownStyleSheet(
                      p: const TextStyle(color: Colors.white),
                      code: TextStyle(color: Colors.white, backgroundColor: Colors.grey[850]),
                      codeblockDecoration: BoxDecoration(
                        color: Colors.grey[850],
                        borderRadius: BorderRadius.circular(8),
                      ),
                    ),
                  )
                : Image.network(
                    message,
                    errorBuilder: (context, error, stackTrace) {
                      return const Text('图片加载失败');
                    },
                  ),
            Padding(
              padding: const EdgeInsets.only(top: 5),
              child: Text(
                messageStruct.createdAt ?? DateTime.now().toString(),
                style: TextStyle(
                  fontSize: 10,
                  color: Colors.white.withOpacity(0.6),
                ),
                textAlign: role == "user" ? TextAlign.right : TextAlign.left,
              ),
            ),
          ],
        ),
      ),
    );
  }

  LinearGradient _getBubbleGradient(String role) {
    return LinearGradient(
      colors: role == 'user'
          ? [const Color(0xFF00C6FF), const Color(0xFF0078FF)] // Blue gradient for the user
          : [const Color(0xFF8983F7), const Color(0xFFA3DAFB)], // Purple gradient for others
      begin: Alignment.topLeft,
      end: Alignment.bottomRight,
    );
  }

  BorderRadius _getBubbleBorderRadius(String role) {
    return role == 'user'
        ? const BorderRadius.only(
            topLeft: Radius.circular(16),
            bottomLeft: Radius.circular(16),
            bottomRight: Radius.circular(16),
          )
        : const BorderRadius.only(
            topRight: Radius.circular(16),
            bottomLeft: Radius.circular(16),
            bottomRight: Radius.circular(16),
          );
  }
}
