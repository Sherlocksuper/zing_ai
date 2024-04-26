// ignore_for_file: constant_identifier_names

class Constant {
  static const String appName = 'AI';
  static const String CURRENT_VERSION = '1.0.1';
  static const String login = 'Login';

  static const String BASE_URL = "http://zing.nat100.top/api";
  static const String SOCKET_URL = "ws://zing.nat100.top/ws";

  ///AUTH API
  static const String LOGIN = '$BASE_URL/auth/login';
  static const String REGISTER = '$BASE_URL/auth/register';
  static const String RESET_PASSWORD = '$BASE_URL/auth/resetpassword';

  ///USER API
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
  static const String SENDFORIMAGE = '$BASE_URL/chat/sendforimage';

  ///Version API
  static const String AllVERSION = '$BASE_URL/version/all';
  static const String LATESTVERSION = '$BASE_URL/version/latest';

  ///Prompt API
  static const String GETPROMPT = '$BASE_URL/prompt/list';
}

enum AIType {
  Text,
  Image,
}
