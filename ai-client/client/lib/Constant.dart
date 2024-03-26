// ignore_for_file: constant_identifier_names

class Constant {
  static const String appName = 'AI';
  static const String CURRENT_VERSION = '1.0.1';
  static const String login = 'Login';

  //
  // static const String BASE_URL = 'http://119.91.193.117:8080/aiapi';
  // static const String SOCKET_URL = 'ws://119.91.193.117:8080/ws';

  static const String BASE_URL = 'http://10.236.169.48:8080/aiapi';
  static const String SOCKET_URL = 'ws://10.236.169.48:8080/ws';

  ///USER API
  static const String LOGIN = '$BASE_URL/user/login';
  static const String REGISTER = '$BASE_URL/user/register';
  static const String USER_INFO = '$BASE_URL/user/find';
  static const String SEND_REGISTER_CODE = '$BASE_URL/user/getemailcode';
  static const String CHECK_REGISTER_CODE = '$BASE_URL/user/checkemailcode';

  ///CHAT API
  static const String StartAChatHAT = '$BASE_URL/chat/start';
  static const String DELETECHAT = '$BASE_URL/chat/delete';
  static const String DELETEALLCHAT = '$BASE_URL/chat/deleteall';
  static const String GETCHATDETAIL = '$BASE_URL/chat/detail';
  static const String GETCHATLIST = '$BASE_URL/chat/list';
  static const String SENDMESSAGE = '$BASE_URL/chat/send';

  ///Version API
  static const String AllVERSION = '$BASE_URL/version/all';
  static const String LATESTVERSION = '$BASE_URL/version/latest';
}
