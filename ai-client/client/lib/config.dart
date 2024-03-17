import 'package:client/Controller/WebSocket.dart';
import 'package:dio/dio.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'Controller/ChatController.dart';

Dio dio = Dio();

@override
void dependencies() {
  Get.lazyPut<ChatController>(() => ChatController());
}

//初始化登录之后的
void afterLogin() {
  WSController.init();
}
