import 'package:client/Controller/chat_controller.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';

import '../../model/chat_struct.dart';

class ChatItem extends StatelessWidget {
  final ChatDetailStruct chat;

  const ChatItem({super.key, required this.chat});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onLongPress: () {
        Get.defaultDialog(
          title: 'Delete Chat',
          content: const Text('Are you sure to delete this chat?'),
          actions: [
            TextButton(
              onPressed: () {
                Get.back();
              },
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: () {
                Get.find<ChatController>().deleteChat(chat.id);
                Get.back();
              },
              child: const Text('Confirm'),
            ),
          ],
        );
      },
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
        child: Row(
          children: [
            Container(
              margin: const EdgeInsets.only(right: 15),
              child: CircleAvatar(
                radius: 30,
                backgroundImage: chat.type == "Text"
                    ? const AssetImage('assets/logo.png')
                    : const AssetImage('assets/beauty_logo.png'),
                backgroundColor: Colors.transparent,
              ),
            ),
            Expanded(
              child: Container(
                padding: const EdgeInsets.all(10),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(15),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.grey.withOpacity(0.2),
                      spreadRadius: 2,
                      blurRadius: 5,
                      offset: const Offset(0, 3),
                    ),
                  ],
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      chat.title,
                      style: const TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                        color: Colors.black87,
                      ),
                    ),
                    const SizedBox(height: 5),
                    Text(
                      chat.messages.isNotEmpty ? chat.messages.last.content : "",
                      style: const TextStyle(
                        fontSize: 16,
                        color: Colors.black54,
                      ),
                      maxLines: 1,
                      overflow: TextOverflow.ellipsis,
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
