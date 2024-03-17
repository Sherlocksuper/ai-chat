import 'package:client/Constant.dart';
import 'package:client/Controller/UserController.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class WSController {
  static bool hasInit = false;
  static late Uri wsUrl;
  static late WebSocketChannel channel;

  static void _judgeNull() {
    if (!hasInit) init();
    print("重新初始化");
  }

  static void init() {
    if (hasInit) return;
    try {
      wsUrl = Uri.parse("${Constant.SOCKET_URL}?userId=${UserController.me.id}");
      channel = IOWebSocketChannel.connect(
        wsUrl,
        pingInterval: const Duration(seconds: 5),
      );

      channel.stream.listen(
        (event) {
          print(event);
        },
        onDone: () {
          hasInit = false;
          print("onDone");
        },
        onError: (e) {
          print(e);
        },
      );
      hasInit = true;
      print("初始化成功");
    } catch (e) {
      print(e);
      hasInit = false;
    }
  }

  //向服务器发送消息
  static void send(String message) {
    _judgeNull();
    channel.sink.add(message);
  }
}
