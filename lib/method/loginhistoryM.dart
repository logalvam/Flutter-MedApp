// To parse this JSON data, do
//
//     final loginHisView = loginHisViewFromJson(jsonString);

import 'dart:convert';
import 'package:intl/intl.dart';

LoginHisView loginHisViewFromJson(String str) =>
    LoginHisView.fromJson(json.decode(str));

String loginHisViewToJson(LoginHisView data) => json.encode(data.toJson());

class LoginHisView {
  List<Hislist> hislist;

  LoginHisView({
    required this.hislist,
  });

  factory LoginHisView.fromJson(Map<String, dynamic> json) => LoginHisView(
        hislist:
            List<Hislist>.from(json["hislist"].map((x) => Hislist.fromJson(x))),
      );

  Map<String, dynamic> toJson() => {
        "hislist": List<dynamic>.from(hislist.map((x) => x.toJson())),
      };
}

class Hislist {
  String userid;
  DateTime date;
  String logintime;
  String logouttime;

  Hislist({
    required this.userid,
    required this.date,
    required this.logintime,
    required this.logouttime,
  });
  DateFormat dateFormat = DateFormat('yyyy-MM-dd');

  factory Hislist.fromJson(Map<String, dynamic> json) => Hislist(
        userid: json["userid"],
        date: DateTime.parse(json["date"]),
        logintime: json["logintime"],
        logouttime: json["logouttime"],
      );

  Map<String, dynamic> toJson() => {
        "userid": userid,
        "date": dateFormat.format(date),
        "logintime": logintime,
        "logouttime": logouttime,
      };
}
