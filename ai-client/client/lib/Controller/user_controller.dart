import 'dart:convert';

import 'package:client/main.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:get/get.dart';
import 'package:url_launcher/url_launcher.dart';
import '../Constant.dart';
import '../config.dart';
import '../pages/Login/login.dart';
import 'file_controller.dart';

///版本检查
class VersionService {
  ///检验版本号是否一致
  Future<void> checkVersion() async {
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

  void updateAppAlert(String updateUrl) {
    Get.defaultDialog(
      title: 'Update Required',
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
      // Rounded corners
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(
            Icons.system_update,
            color: Colors.cyanAccent,
            size: 60,
          ),
          SizedBox(height: 16),
          Text(
            'Current version ${Constant.CURRENT_VERSION} is not available. Please update.',
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 14,
              color: Colors.grey[100],
            ),
          ),
          SizedBox(height: 16),
          Text(
            'Would you like to update now?',
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 14,
              color: Colors.grey[100],
            ),
          ),
        ],
      ),
      textConfirm: 'YES',
      textCancel: 'NO',
      confirmTextColor: Colors.white,
      buttonColor: Colors.cyanAccent,
      // Confirm button color
      cancelTextColor: Colors.grey[300],
      onConfirm: () {
        if (updateUrl.isEmpty) {
          updateUrl = "https://www.baidu.com";
        }
        launchUrl(Uri.parse(updateUrl));
      },
      onCancel: () {
        SystemNavigator.pop();
      },
    );
  }
}

///本地用户信息
class LocalAuthService {
  static const String userFile = "user.txt";

  ///TODO:写入userFile信息
  Future<void> writeUserFile(User me) async {
    print("正在写入userFile信息 , location : user_controller.dart ,writeUserFile");
    await FileController.writeToFile(userFile, json.encode(User.toJson(me)));
  }

  ///TODO:读取userFile信息
  Future<String> readUserFile() async {
    var data = await FileController.readFromFile(userFile);
    if (data == "") return "";
    return data;
  }

  ///TODO:清空userFile信息
  Future<void> clearUserFile() async {
    await FileController.clearFileData(userFile);
  }
}

///用户控制器
class UserController {
  static VersionService versionService = VersionService();
  static LocalAuthService authService = LocalAuthService();

  static String userFile = "user.txt";
  static User me = User(id: 0, name: '', password: '', token: '', email: '');

  ///检查登录
  static Future<void> checkLogin() async {
    print('checkLogin');
    var data = await authService.readUserFile();
    if (data == "") {
      Get.off(() => LoginRegisterPage());
    }
    me = User.fromJson(json.decode(data));
    Get.off(() => HomeTab());
    //
    print("checklogin 完成，flutter_native_splash.remove(); location : user_controller.dart ,checkLogin");
    FlutterNativeSplash.remove();
    versionService.checkVersion();
  }

  ///登录
  static Future<bool> login(String name, String password) async {
    print('登录中 , location : user_controller.dart ,login');
    var response = await dio.post(Constant.LOGIN, data: {'name': name, 'password': password});
    print("登录返回信息:$response");
    if (response.data["code"] == 200) {
      me = User(
        id: response.data["data"]["id"],
        name: response.data["data"]["name"],
        password: response.data["data"]["password"],
        token: response.data["data"]["token"],
        email: response.data["data"]["email"],
      );
      await authService.writeUserFile(me);
      Get.off(() => HomeTab());
      versionService.checkVersion();
      return true;
    } else {
      EasyLoading.showError(response.data["message"]);
      return false;
    }
  }

  ///注册
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

  ///退出登录
  static Future<bool> logout() async {
    Get.offAll(() => LoginRegisterPage());
    await authService.clearUserFile();
    return true;
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

  static User fromJson(decode) {
    return User(
      id: decode["id"],
      name: decode["name"],
      password: decode["password"],
      token: decode["token"] ?? "",
      email: decode["email"] ?? "",
    );
  }

  static Map<String, dynamic> toJson(User user) {
    return {
      "id": user.id,
      "name": user.name,
      "password": user.password,
      "token": user.token,
      "email": user.email,
    };
  }
}
