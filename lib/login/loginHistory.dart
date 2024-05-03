import 'package:flutter/material.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/method/loginhistoryM.dart';

import '../apimethods/logHisViewM.dart';

class LoginHistoryView extends StatefulWidget {
  const LoginHistoryView({super.key});

  @override
  State<LoginHistoryView> createState() => _LoginHistoryViewState();
}

LoginHisViewA logView = LoginHisViewA();
var logRecodrs;
var dumm = "";
LoginHisView? HisView;

class _LoginHistoryViewState extends State<LoginHistoryView> {
  List? login;
  int? lengthofarray;
  bool isloading = true;
  @override
  void setState(VoidCallback fn) {
    // TODO: implement setState
    print("object");

    super.setState(fn);
  }

  @override
  void initState() {
    // TODO: implement initState
    fethcLogHis();
    super.initState();
  }

  fethcLogHis() async {
    var History = await logView.LoginHisViewApi();
    // print(History.toString());
    setState(() {
      HisView = logView.model;
      dumm = "false";
    });
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    TextStyle fonts = TextStyle(
        fontSize: screenWidth * 0.03,
        fontFamily: "Helvetica",
        fontWeight: FontWeight.bold);
    return Scaffold(
      body: ListView(
        scrollDirection: Axis.vertical,
        children: [
          Center(
              child: dumm.isEmpty
                  ? Center(
                      child: CircularProgressIndicator(),
                    )
                  : Container(
                      color: AppColors.bgColor,
                      child: Row(
                        children: [
                          Column(
                            children: [
                              Container(
                                width: screenWidth,
                                child: DataTable(
                                    columnSpacing: 4,
                                    columns: <DataColumn>[
                                      DataColumn(
                                          tooltip: "Name",
                                          label: Text(
                                            " Name",
                                            style: TextStyle(
                                                // fontSize: screenWidth * 0.04,
                                                fontStyle: FontStyle.italic,
                                                fontWeight: FontWeight.bold),
                                          )),
                                      DataColumn(
                                          label: Text(
                                        "Date",
                                        style: TextStyle(
                                            // fontSize: screenWidth * 0.04,
                                            fontStyle: FontStyle.italic,
                                            fontWeight: FontWeight.bold),
                                      )),
                                      DataColumn(
                                          label: Text(
                                        "Login",
                                        style: TextStyle(
                                            fontStyle: FontStyle.italic,
                                            // fontSize: screenWidth * 0.04,
                                            fontWeight: FontWeight.bold),
                                      )),
                                      DataColumn(
                                          label: Text(
                                        "Logout",
                                        style: TextStyle(
                                            // fontSize: screenWidth * 0.04,
                                            fontStyle: FontStyle.italic,
                                            fontWeight: FontWeight.bold),
                                      )),
                                    ],
                                    rows: List<DataRow>.generate(
                                        HisView!.hislist.length,
                                        (index) => DataRow(cells: <DataCell>[
                                              DataCell(
                                                Text(
                                                  HisView!
                                                      .hislist[index].userid,
                                                  style: fonts,
                                                ),
                                              ),
                                              DataCell(
                                                Text(
                                                  HisView!.hislist[index].date
                                                      .toString()
                                                      .replaceAll(
                                                          "00:00:00.000", ""),
                                                  style: fonts,
                                                ),
                                              ),
                                              DataCell(Text(
                                                  HisView!
                                                      .hislist[index].logintime,
                                                  style: fonts)),
                                              DataCell(Text(
                                                HisView!
                                                    .hislist[index].logouttime,
                                                style: fonts,
                                              )),
                                            ]))),
                              ),
                            ],
                          ),
                        ],
                      ),
                    )),
        ],
      ),
    );
  }
}
