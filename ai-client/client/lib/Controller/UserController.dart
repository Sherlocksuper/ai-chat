import 'package:client/main.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:url_launcher/url_launcher.dart';
import '../Constant.dart';
import '../config.dart';
import '../pages/Login/Login.dart';

class UserController {
  static User me = User(id: 0, name: '', password: '', token: '', email: '');

  //登录
  static Future<bool> login(String name, String password) async {
    print('login');
    var response = await dio.post(Constant.LOGIN, data: {'name': name, 'password': password});
    print(response);
    if (response.data["code"] == 200) {
      me = User(
        id: response.data["data"]["id"],
        name: response.data["data"]["name"],
        password: response.data["data"]["password"],
        token: response.data["data"]["token"],
        email: response.data["data"]["email"],
      );
      Get.off(() => HomeTab());
      checkVersion();
      return true;
    } else {
      EasyLoading.showError(response.data["message"]);
      return false;
    }
  }

  //注册
  static Future<bool> register(String name, String password) async {
    var response = await dio.post(Constant.REGISTER, data: {'name': name, 'password': password});
    print(response);
    if (response.data["code"] == 200) {
      EasyLoading.showSuccess('注册成功,请返回登录');
      return true;
    } else {
      EasyLoading.showError('注册失败,${response.data["message"]}');
      return false;
    }
  }

  //退出登录
  static Future<bool> logout() async {
    Get.offAll(() => LoginRegisterPage());
    return true;
  }

  //检验版本号是否一致
  static Future<void> checkVersion() async {
    var response = await dio.get(Constant.AllVERSION);
    if (response.data["code"] == 200) {
      print(response);
      var res = (response.data["data"] as List).firstWhere((element) => element["version"] == Constant.CURRENT_VERSION);
      if (res["enable"]) return;
      launchUrl(Uri.parse(res["downloadUrl"]));
      updateAppAlert(res["downloadUrl"]);
    } else {
      EasyLoading.showError('版本检查失败,请检查网络');
    }
  }

  static void updateAppAlert(String updateUrl) {
    Get.defaultDialog(
      title: '提示,当前版本${Constant.CURRENT_VERSION}不可用,请更新',
      middleText: '是否更新',
      textConfirm: '是',
      textCancel: '否',
      confirmTextColor: Colors.white,
      onConfirm: () {
        if (updateUrl == "") updateUrl == "www.baidu.com";
        launchUrl(Uri.parse(updateUrl));
      },
      onCancel: () {
        SystemNavigator.pop();
      },
    );
  }
}

class User {
  int id;
  String name;
  String password;
  String token;
  String email;

  User({
    required this.id,
    required this.name,
    required this.password,
    required this.token,
    required this.email,
  });
}
