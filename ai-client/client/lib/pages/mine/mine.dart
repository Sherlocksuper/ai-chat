import 'package:client/Controller/web_socket.dart';
import 'package:client/pages/mine/prompts.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
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
            ListTile(
              leading: const Icon(Icons.notifications),
              title: const Text('Prompts'),
              onTap: () {
                Get.to(() =>  PromptList());
              },
            ),
            //黑夜模式
            ListTile(
              onTap: () {
                //切换黑夜模式
                Get.changeTheme(Get.isDarkMode ? ThemeData.light() : ThemeData.dark());
              },
              leading: const Icon(Icons.nightlight_round),
              title: const Text('黑夜模式'),
            ),
            //当前版本
            ListTile(
              leading: const Icon(Icons.info),
              title: const Text('当前版本   点击检查更新'),
              trailing: const Text(Constant.CURRENT_VERSION),
              onTap: () async {
                await UserController.versionService.checkLatestVersion();
              },
              // Rounded corners
            ),
            ListTile(
              leading: const Icon(Icons.logout, color: Colors.red),
              title: const Text('退出登录'),
              onTap: () {
                UserController.logout();
              },
            ),
          ],
        ),
      ),
    );
  }
}
