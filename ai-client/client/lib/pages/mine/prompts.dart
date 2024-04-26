import 'package:async/async.dart';
import 'package:client/Controller/setting_ontroller.dart';
import 'package:easy_refresh/easy_refresh.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:get/get_state_manager/src/simple/get_state.dart';

import '../../Constant.dart';
import '../../config.dart';

class PromptList extends StatelessWidget {
  PromptList({super.key});

  List<Prompt> prompts = [];
  int selectedIndex = -1;
  Function getData = tgetPrompts();

  @override
  Widget build(BuildContext context) {
    if (!Get.isRegistered<SettingController>()) Get.lazyPut<SettingController>(() => SettingController());

    return FutureBuilder(
      future: () async {
        prompts = prompts + await getData();
      }(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        }
        return Scaffold(
          appBar: AppBar(
            title: const Text('Prompts'),
          ),
          body: EasyRefresh(
            onLoad: () async {
              prompts = prompts + await getData();
              Get.find<SettingController>().update(['promptList']);
            },
            child: GetBuilder<SettingController>(
              id: 'promptList',
              builder: (logic) {
                return ListView.builder(
                  itemCount: prompts.length,
                  itemBuilder: (context, index) {
                    return GestureDetector(
                      onTap: () {
                        selectedIndex = selectedIndex == index ? -1 : index;
                        logic.update(['promptList']);
                      },
                      child: PromptItem(
                        isSelected: selectedIndex == index,
                        prompt: prompts[index],
                      ),
                    );
                  },
                );
              },
            ),
          ),
        );
      },
    );
  }
}

class PromptItem extends StatelessWidget {
  PromptItem({
    super.key,
    required this.isSelected,
    required this.prompt,
  });

  bool isSelected = false;
  final Prompt prompt;

  @override
  Widget build(BuildContext context) {
    return Card(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
      ),
      elevation: 4,
      margin: const EdgeInsets.all(10),
      child: Padding(
        padding: const EdgeInsets.all(10),
        child: Row(
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    prompt.title,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                    ),
                  ),
                  const SizedBox(height: 5),
                  AnimatedSize(
                    duration: const Duration(milliseconds: 300),
                    child: Column(
                      children: [
                        Text(
                          prompt.content,
                          style: const TextStyle(fontSize: 14),
                          overflow: TextOverflow.clip,
                          maxLines: isSelected ? 99 : 3,
                        ),
                        const Divider(),
                        Text(
                          'Function: ${prompt.function}',
                          style: const TextStyle(
                            color: Colors.grey,
                            fontSize: 12,
                          ),
                          maxLines: isSelected ? 99 : 1,
                        ),
                      ],
                    ),
                  )
                ],
              ),
            ),
            IconButton(
              icon: const Icon(Icons.copy),
              onPressed: () {
                Clipboard.setData(ClipboardData(text: prompt.content));
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(
                    content: Text('Prompt copied to clipboard!'),
                  ),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}

class Prompt {
  int id;
  String title;
  String content;
  String function;

  Prompt({
    required this.id,
    required this.title,
    required this.function,
    required this.content,
  });

  //fromJson
  factory Prompt.fromJson(Map<String, dynamic> json) {
    return Prompt(
      id: json['id'] ?? json['ID'],
      title: json['title'],
      content: json['content'],
      function: json['function'] ?? json['title'],
    );
  }
}

//定义一个返回值为函数的函数
Function tgetPrompts() {
  int offset = 0;

  return () async {
    print('offset: $offset');
    var response = await dio.get("${Constant.GETPROMPT}?offset=$offset");
    if (response.data['code'] == 200) {
      print(response.data['data']);
      List<Prompt> prompts = [];
      for (var item in response.data['data']) {
        prompts.add(Prompt.fromJson(item));
      }
      offset++;
      return prompts;
    } else {}
    return [];
  };
}
