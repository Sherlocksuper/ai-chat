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
      updateAppAlert(res["downloadUrl"], true);
    } else {
      EasyLoading.showError('版本检查失败,请检查网络');
    }
  }

  ///检查最新版本
  Future<void> checkLatestVersion() async {
    await EasyLoading.show(status: '检查更新中...');
    var response = await dio.get(Constant.LATESTVERSION);
    await EasyLoading.dismiss();

    if (response.data["code"] == 200) {
      print(response);
      var res = response.data["data"];
      if (res["version"] == Constant.CURRENT_VERSION) {
        EasyLoading.showSuccess('当前已是最新版本');
      } else {
        updateAppAlert(res["downloadUrl"], false);
      }
    } else {
      EasyLoading.showError('版本检查失败,请检查网络');
    }
  }

  ///更新app弹窗
  void updateAppAlert(String updateUrl, bool required) {
    Get.defaultDialog(
      title: required ? 'Update Required' : 'Update Available',
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
      barrierDismissible: !required,
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          const Icon(
            Icons.system_update,
            color: Colors.cyanAccent,
            size: 60,
          ),
          const SizedBox(height: 16),
          Text(
            'Current version ${Constant.CURRENT_VERSION} is not available. Please update.',
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 14,
              color: Colors.grey[100],
            ),
          ),
          const SizedBox(height: 16),
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
      cancelTextColor: Colors.grey[300],
      onConfirm: () {
        if (updateUrl.isEmpty) {
          updateUrl = "https://www.baidu.com";
        }
        launchUrl(Uri.parse(updateUrl));
      },
      onCancel: () {
        if (required) {
          SystemNavigator.pop();
        }
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
    var data = await authService.readUserFile();
    if (data != "") {
      try {
        me = User.fromJson(json.decode(data));
        Get.off(() => HomeTab());
      } catch (e) {
        authService.clearUserFile();
      }
    }

    print("checklogin 完成，flutter_native_splash.remove(); location : user_controller.dart ,checkLogin");
    FlutterNativeSplash.remove();
    versionService.checkVersion();
  }

  ///登录
  static Future<bool> login(String password, String email) async {
    print('登录中 , location : user_controller.dart ,login');
    var response = await dio.post(Constant.LOGIN, data: {'password': password, 'email': email});
    print("登录返回信息:$response");
    if (response.data["code"] == 200) {
      me = User(
        id: response.data["data"]["id"] ?? response.data["data"]["ID"],
        name: response.data["data"]["name"],
        password: response.data["data"]["password"],
        token: response.data["data"]["token"],
        email: response.data["data"]["email"],
      );
      await authService.writeUserFile(me);
      Get.off(() => HomeTab());
      return true;
    } else {
      EasyLoading.showError(response.data["message"]);
      return false;
    }
  }

  ///注册
  static Future<bool> register(String name, String password, String email) async {
    if (name.isEmpty || password.isEmpty || email.isEmpty) {
      EasyLoading.showError('请填写完整信息');
      return false;
    }
    var response = await dio.post(Constant.REGISTER, data: {'name': name, 'password': password, 'email': email});
    print(response);
    if (response.data["code"] == 200) {
      EasyLoading.showSuccess('注册成功,请返回登录');
      return true;
    } else {
      EasyLoading.showError('注册失败,${response.data["message"]}');
      return false;
    }
  }

  ///发送验证码
  static Future<bool> sendVerificationCode(String email) async {
    if (email.isEmpty) {
      EasyLoading.showError('请填写邮箱');
      return false;
    }
    var response = await dio.get(Constant.SEND_REGISTER_CODE, queryParameters: {'email': email});
    print(response);
    if (response.data["code"] == 200) {
      EasyLoading.showSuccess('发送成功,请查收邮箱');
      return true;
    } else {
      EasyLoading.showError('发送失败,${response.data["message"]}');
      return false;
    }
  }

  ///检查验证码
  static Future<bool> checkVerificationCode(String email, String code) async {
    if (email.isEmpty || code.isEmpty) {
      EasyLoading.showError('电子邮箱与验证码不能为空');
      return false;
    }

    var response = await dio.get(Constant.CHECK_REGISTER_CODE, queryParameters: {'email': email, 'code': code});
    print(response);
    if (response.data["code"] == 200) {
      EasyLoading.showSuccess('验证成功');
      return true;
    } else {
      EasyLoading.showError('验证失败,${response.data["message"]}');
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

///用户信息
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
