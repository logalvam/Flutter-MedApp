// To parse this JSON data, do
//
//     final autentication = autenticationFromJson(jsonString);

import 'dart:convert';

Autentication autenticationFromJson(String str) =>
    Autentication.fromJson(json.decode(str));

String autenticationToJson(Autentication data) => json.encode(data.toJson());

class Autentication {
  String role;
  String status;
  String errMeg;

  Autentication({
    required this.role,
    required this.status,
    required this.errMeg,
  });

  factory Autentication.fromJson(Map<String, dynamic> json) => Autentication(
        role: json["role"],
        status: json["status"],
        errMeg: json["errMeg"],
      );

  Map<String, dynamic> toJson() => {
        "role": role,
        "status": status,
        "errMeg": errMeg,
      };
}
