// To parse this JSON data, do
//
//     final todayBsales = todayBsalesFromJson(jsonString);

import 'dart:convert';

TodayBsales todayBsalesFromJson(String str) =>
    TodayBsales.fromJson(json.decode(str));

String todayBsalesToJson(TodayBsales data) => json.encode(data.toJson());

class TodayBsales {
  int sales;
  String date;

  TodayBsales({
    required this.sales,
    required this.date,
  });

  factory TodayBsales.fromJson(Map<String, dynamic> json) => TodayBsales(
        sales: json["sales"],
        date: json["date"],
      );

  Map<String, dynamic> toJson() => {
        "sales": sales,
        "date": date,
      };
}
