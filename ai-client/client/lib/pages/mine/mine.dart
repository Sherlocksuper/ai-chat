import 'package:client/Controller/web_socket.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';

import '../../Constant.dart';
import '../../Controller/user_controller.dart';

class Mine extends StatelessWidget {
  Mine({super.key});

  bool isDarkMode = Get.isDarkMode;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text('我的'),
        ),
        body: Column(
          children: [
            //我的名字
            ListTile(
              leading: const Icon(Icons.person),
              title: Text("UserName:${UserController.me.name}"),
            ),
            //黑夜模式
            ListTile(
              onTap: () {
                isDarkMode = !isDarkMode;
                if (isDarkMode) {
                  Get.changeTheme(ThemeData.dark());
                } else {
                  Get.changeTheme(ThemeData.light());
                }
              },
              leading: const Icon(Icons.nightlight_round),
              title: const Text('黑夜模式'),
            ),
            ListTile(
              leading: const Icon(Icons.logout, color: Colors.red),
              title: const Text('退出登录'),
              onTap: () {
                UserController.logout();
              },
            ),
            //发送一个websocket
            ListTile(
              leading: const Icon(Icons.send),
              title: const Text('发送一个websocket'),
              onTap: () {
                WSController.send("客户端尝试");
              },
            ),

            //当前版本
            const ListTile(
              leading: Icon(Icons.info),
              title: Text('当前版本'),
              trailing: Text(Constant.CURRENT_VERSION),
            ),
          ],
        ),
      ),
    );
  }
}
