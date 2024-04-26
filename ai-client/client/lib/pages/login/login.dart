import 'package:client/pages/login/reset.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import '../../Controller/user_controller.dart';

class LoginRegisterPage extends StatefulWidget {
  @override
  _LoginRegisterPageState createState() => _LoginRegisterPageState();
}

class _LoginRegisterPageState extends State<LoginRegisterPage> {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final TextEditingController _emailController = TextEditingController(); // 新增
  final TextEditingController _verificationCodeController = TextEditingController(); // 新增
  bool _isLogin = true;

  DateTime? lastSendTime;

  @override
  void initState() {
    super.initState();
    UserController.checkLogin();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: true,
      appBar: AppBar(
        title: Text(_isLogin ? 'Login' : 'Register'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Center(
          child: SingleChildScrollView(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                if (!_isLogin) BuildTextInput(controller: _usernameController, label: 'Username', icon: Icons.person),
                const SizedBox(height: 20.0),
                BuildTextInput(controller: _emailController, label: 'Email', icon: Icons.email), // 新增
                const SizedBox(height: 20.0),
                BuildTextInput(controller: _passwordController, label: 'Password', icon: Icons.lock, isPassword: true),
                const SizedBox(height: 20.0),
                if (!_isLogin) ...[
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Expanded(
                        child: BuildTextInput(
                            controller: _verificationCodeController, label: 'Verification Code', icon: Icons.lock),
                      ),
                      const SizedBox(width: 20.0),
                      ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          foregroundColor: Colors.white, backgroundColor: Colors.blue, // Text color
                        ),
                        onPressed: () async {
                          if (lastSendTime != null &&
                              DateTime.now().difference(lastSendTime!) < const Duration(seconds: 60)) {
                            await EasyLoading.showError('请等待60秒后再次发送');
                            return;
                          }
                          bool sendSuccess = await UserController.sendVerificationCode(_emailController.text);
                          if (sendSuccess) lastSendTime = DateTime.now();
                        },
                        child: const Text('发送验证码'), // 新增
                      ),
                    ],
                  ),
                  const SizedBox(height: 20.0),
                ],
                ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.white, backgroundColor: Colors.grey[800], // Text color
                  ),
                  onPressed: () async {
                    String username = _usernameController.text;
                    String password = _passwordController.text;
                    String email = _emailController.text; // 新增
                    if (_isLogin) {
                      UserController.login(password, email);
                    } else {
                      String verificationCode = _verificationCodeController.text; // 新增
                      bool checkVerSuccess = await UserController.checkVerificationCode(email, verificationCode); // 新增
                      if (checkVerSuccess) {
                        UserController.register(username, password, email);
                      }
                    }
                  },
                  child: Text(_isLogin ? 'Login' : 'Register'),
                ),
                TextButton(
                  style: TextButton.styleFrom(
                    foregroundColor: Colors.grey[500], // Grey text button
                  ),
                  onPressed: () {
                    setState(() {
                      clearAll();
                      _isLogin = !_isLogin;
                    });
                  },
                  child: Text(_isLogin ? 'Create an account' : 'Already have an account? Login'),
                ),
                TextButton(
                  style: TextButton.styleFrom(
                    foregroundColor: Colors.grey[500], // Grey text button
                  ),
                  onPressed: () {
                    Get.to(ResetPassword());
                  },
                  child: Text('Forgot password?'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  void clearAll() {
    _usernameController.clear();
    _passwordController.clear();
    _emailController.clear();
    _verificationCodeController.clear();
  }
}

class BuildTextInput extends StatefulWidget {
  final TextEditingController controller;
  final String label;
  final IconData icon;
  final bool isPassword;

  const BuildTextInput(
      {super.key, required this.controller, required this.label, required this.icon, this.isPassword = false});

  @override
  State<BuildTextInput> createState() => _BuildTextInputState();
}

class _BuildTextInputState extends State<BuildTextInput> {
  late bool unShow = widget.isPassword;

  TextInputType judgeType(String label) {
    switch (label) {
      case 'Email':
        return TextInputType.emailAddress;
      case 'Verification Code':
        return TextInputType.number;
      case 'Password':
        return TextInputType.text;
      default:
        return TextInputType.text;
    }
  }

  @override
  Widget build(BuildContext context) {
    return TextField(
      cursorColor: Colors.grey,
      obscureText: unShow,
      controller: widget.controller,
      decoration: InputDecoration(
        label: Text(widget.label, style: const TextStyle(color: Colors.black)),
        prefixIcon: Icon(widget.icon),
        enabledBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.grey),
        ),
        disabledBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.grey),
        ),
        focusedBorder: const OutlineInputBorder(
          borderSide: BorderSide(color: Colors.black),
        ),
        focusColor: Colors.red,
        //pre 眼睛
        suffixIcon: widget.isPassword
            ? IconButton(
                icon: const Icon(Icons.remove_red_eye),
                onPressed: () {
                  setState(() {
                    unShow = !unShow;
                  });
                },
              )
            : null,
      ),
      style: const TextStyle(),
      onChanged: (value) {
        widget.controller.text = value;
      },
      keyboardType: judgeType(widget.label),
    );
  }
}
