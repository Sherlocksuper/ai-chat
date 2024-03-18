// ignore_for_file: constant_identifier_names

class Constant {
  static const String appName = 'AI';
  static const String CURRENT_VERSION = '1.0.1';
  static const String login = 'Login';

  static const String BASE_URL = 'http://192.168.0.188:8080';
  static const String SOCKET_URL = 'ws://192.168.0.188:8080/ws';

  ///USER API
  static const String LOGIN = '$BASE_URL/user/login';
  static const String REGISTER = '$BASE_URL/user/register';
  static const String USER_INFO = '$BASE_URL/user/find';

  ///CHAT API
  static const String StartAChatHAT = '$BASE_URL/chat/start';
  static const String DELETECHAT = '$BASE_URL/chat/delete';
  static const String DELETEALLCHAT = '$BASE_URL/chat/deleteall';
  static const String GETCHATDETAIL = '$BASE_URL/chat/detail';
  static const String GETCHATLIST = '$BASE_URL/chat/list';
  static const String SENDMESSAGE = '$BASE_URL/chat/send';

  ///Version API
  static const String AllVERSION = '$BASE_URL/version/all';
}
