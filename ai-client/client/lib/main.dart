import 'package:client/Controller/user_controller.dart';
import 'package:client/KeepAliveWrapper.dart';
import 'package:client/config.dart';
import 'package:client/pages/Chat/chat.dart';
import 'package:client/pages/Login/login.dart';
import 'package:client/pages/Mine/mine.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:get/get.dart';

void main() {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  afterIn();
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'GetX Tutorial',
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.blue,
          secondary: Colors.blueAccent,
          primary: Colors.blueAccent,
        ),
        tabBarTheme: const TabBarTheme(
          labelColor: Colors.blue,
          unselectedLabelColor: Colors.grey,
        ),
        cardColor: Colors.white,
      ),
      home: LoginRegisterPage(),
      builder: EasyLoading.init(),
      debugShowCheckedModeBanner: false,
    );
  }
}

class HomeTab extends StatelessWidget {
  final List<HomePageType> homePageType = [
    HomePageType('聊天', Icons.home, const ChatPage()),
    HomePageType('我的', Icons.person, Mine()),
  ];

  HomeTab({super.key});

  @override
  Widget build(BuildContext context) {
    afterLogin();
    return DefaultTabController(
      length: homePageType.length,
      child: Scaffold(
        body: TabBarView(
          children: List.generate(
            homePageType.length,
            (index) => KeepAliveWrapper(child: homePageType[index].page),
          ),
        ),
        bottomNavigationBar: TabBar(
          tabs: List.generate(
            homePageType.length,
            (index) => Tab(
              icon: Icon(homePageType[index].icon),
              text: homePageType[index].title,
            ),
          ),
        ),
      ),
    );
  }
}

class HomePageType {
  final IconData icon;
  final String title;
  final Widget page;

  HomePageType(this.title, this.icon, this.page);
}
