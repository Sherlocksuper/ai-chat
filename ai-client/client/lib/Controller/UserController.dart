import 'package:client/main.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';

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
      Get.to(() => HomeTab());
      afterLogin();
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
    Get.to(() => LoginRegisterPage());
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
}
