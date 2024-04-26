//忘记密码后重设密码页，通过邮箱验证重设密码
import 'package:client/Controller/user_controller.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'login.dart';

class ResetPassword extends StatelessWidget {
  final emailController = TextEditingController();
  final codeController = TextEditingController();
  final passwordController = TextEditingController();
  final passwordConfirmController = TextEditingController();

  DateTime? lastSendTime;

  ResetPassword({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Reset Password'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Center(
          child: SingleChildScrollView(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                BuildTextInput(controller: emailController, label: 'Email', icon: Icons.email),
                const SizedBox(height: 20.0),
                Row(
                  children: [
                    Expanded(child: BuildTextInput(controller: codeController, label: 'Code', icon: Icons.lock)),
                    const SizedBox(width: 10.0),
                    ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        foregroundColor: Colors.white, backgroundColor: Colors.blue, // Text color
                      ),
                      onPressed: () async {
                        if (lastSendTime != null &&
                            DateTime.now().difference(lastSendTime!) < const Duration(seconds: 60)) {
                          EasyLoading.showToast('Please wait for 60 seconds');
                          return;
                        }

                        if (emailController.text.isEmpty || !emailController.text.contains('@')) {
                          EasyLoading.showToast('Please input correct email');
                          return;
                        }

                        await EasyLoading.show(status: 'Sending code...');
                        bool sendSuccess = await UserController.sendVerificationCode(emailController.text);
                        if (sendSuccess) lastSendTime = DateTime.now();
                        await EasyLoading.dismiss();
                      },
                      child: const Text('Send Code'), // 新增
                    ),
                  ],
                ),
                const SizedBox(height: 20.0),
                BuildTextInput(controller: passwordController, label: 'Password', icon: Icons.lock, isPassword: true),
                const SizedBox(height: 20.0),
                BuildTextInput(
                    controller: passwordConfirmController,
                    label: 'Confirm Password',
                    icon: Icons.lock,
                    isPassword: true),
                const SizedBox(height: 20.0),
                ElevatedButton(
                  onPressed: () async {
                    bool check = await UserController.checkVerificationCode(emailController.text, codeController.text);
                    if (!check) {
                      EasyLoading.showToast('Verification code is wrong');
                      return;
                    }
                    if (checkResetPassword()) {
                      UserController.resetPassword(emailController.text, passwordController.text);
                    }
                  },
                  child: const Text('Reset Password'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  //检查重设密码是否填写了所有信息
  bool checkResetPassword() {
    if (emailController.text.isEmpty || !emailController.text.contains('@')) {
      EasyLoading.showToast('Please input correct email');
      return false;
    }
    if (codeController.text.isEmpty) {
      EasyLoading.showToast('Please input code');
      return false;
    }
    if (passwordController.text.isEmpty) {
      EasyLoading.showToast('Please input password');
      return false;
    }
    if (passwordConfirmController.text.isEmpty) {
      EasyLoading.showToast('Please input confirm password');
      return false;
    }
    return true;
  }
}
