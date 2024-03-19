import 'package:client/Controller/web_socket.dart';
import 'package:dio/dio.dart';
import 'package:get/get.dart';
import 'Controller/chat_controller.dart';

Dio dio = Dio();

//初始化登录之后的
void afterLogin() {
  WSController.init();
  Get.lazyPut<ChatController>(() => ChatController());
}
