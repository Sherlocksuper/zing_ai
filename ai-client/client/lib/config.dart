import 'package:client/Controller/setting_ontroller.dart';
import 'package:client/Controller/web_socket.dart';
import 'package:dio/dio.dart';
import 'package:get/get.dart';
import 'Controller/chat_controller.dart';

Dio dio = Dio();

void afterIn() {
  Get.lazyPut<SettingController>(() => SettingController());
  Get.lazyPut<ChatController>(() => ChatController());
}

//初始化登录之后的
void afterLogin() {
  WSController.init();
}
