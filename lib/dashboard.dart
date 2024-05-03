import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:medapp/apimethods/todaySalesM.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/login/loginPage.dart';
import 'package:medapp/stock/stockView.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'apimethods/inventryvalye.dart';
import 'apimethods/logoutHistoryM.dart';
import 'bill/billentry.dart';
import 'login/addUser.dart';
import 'login/loginHistory.dart';
import 'salesreposrt.dart';
import 'stock/stockEntry.dart';

class DashBoard extends StatefulWidget {
  const DashBoard({super.key});

  @override
  State<DashBoard> createState() => _DashBoardState();
}

LogoutHisInsertA loutHis = LogoutHisInsertA();
String email = "";
TodaySalesA tSales = TodaySalesA();
var sales;
String? value;
String userrole = "";
var screenWidth;

class _DashBoardState extends State<DashBoard> with TickerProviderStateMixin {
  // late TabController tabController1;
  List<Widget> tabs = [];
  List<Widget> dashboardbody = [];
  String UserBar = "";
  @override
  void initState() {
    super.initState();
    initializeDashboard();
    fetchTodasales();
    setState(() {});
  }

  Future<dynamic> fetchTodasales() async {
    sales = await tSales.TodaySalesApi(email);
    value = sales["sales"].toString();
    print("value :${value.runtimeType}");
    setState(() {});
    return;
  }

  void initializeDashboard() async {
    SharedPreferences sref = await SharedPreferences.getInstance();
    userrole = await sref.getString("role")!;
    email = await sref.getString("user")!;
    print("role: $userrole");
    assignDashboard();
    // tabController1 =TabController(length: tabs.length, vsync: this, initialIndex: 0);
    setState(() {});
  }

  void assignDashboard() {
    // Update tabs and dashboardbody based on userrole
    // This is just an example, adjust according to your needs
    if (userrole == 'Biller') {
      UserBar = userrole;
      tabs = [
        Tab(
            child: Text(
          "Dash Board",
          style: TextStyle(fontSize: screenWidth * 0.03),
        )),
        Tab(
          child: Text(
            "Bill Entry",
            style: TextStyle(fontSize: screenWidth * 0.03),
          ),
        ),
        Tab(
            child: Text(
          "Stock View",
          style: TextStyle(fontSize: screenWidth * 0.03),
        )),
      ];
      ;
      dashboardbody = [
        billerDash(),
        Billentry(),
        StockViewd(),
      ];
      setState(() {});
    } else if (userrole == "Manager") {
      UserBar = userrole;

      tabs = [
        Tab(child: Text("Dash Board")),
        Tab(child: Text("Stock Entry")),
        Tab(child: Text("Sales Report")),
        Tab(child: Text("Stock View")),
      ];
      ;
      dashboardbody = [
        managerDash(),
        StockEntry(),
        SalesReport(),
        StockViewd(),
      ];
      setState(() {});
    } else if (userrole == 'Inventry') {
      UserBar = userrole;

      tabs = [
        Tab(child: Text("Dash Board")),
        Tab(child: Text("Stock Entry")),
        Tab(child: Text("Stock View")),
      ];
      ;
      dashboardbody = [
        AdminnInventry(),
        StockEntry(),
        StockViewd(),
      ];
      setState(() {});
    } else if (userrole == 'Admin') {
      UserBar = userrole;

      // Handle other roles similarly
      tabs = [
        Tab(
            child: Text(
          "Dash Board",
          style: TextStyle(fontSize: screenWidth * 0.045),
        )),
        Tab(
            child: Text(
          "Add User",
          style: TextStyle(fontSize: screenWidth * 0.045),
        )),
        Tab(
            child: Text(
          "Login History",
          style: TextStyle(fontSize: screenWidth * 0.045),
        )),
      ];
      ;
      dashboardbody = [
        AdminnInventry(),
        UserAddView(),
        LoginHistoryView(),
      ];
      setState(() {}); // Call setState to update the UI
    }
  }

  @override
  Widget build(BuildContext context) {
    screenWidth = MediaQuery.of(context).size.width;
    return DefaultTabController(
      length: tabs.length,
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: Text(UserBar ?? 'Welcome'),
          actions: [
            IconButton(
              icon: Icon(Icons.logout),
              onPressed: () async {
                SharedPreferences sref = await SharedPreferences.getInstance();
                var resp = loutHis.LogoutHisInsertApi();
                print("logout $resp");
                sref.setString("role", "");
                sref.setString("user", "");
                Navigator.pop(
                  context,
                  MaterialPageRoute(
                    builder: (context) => Loginpage(),
                  ),
                );
              },
            ),
          ],
          bottom: TabBar(
            tabs: tabs,
          ),
        ),
        body: email.isEmpty
            ? Center(
                child: CircularProgressIndicator(),
              )
            : TabBarView(
                children: dashboardbody,
              ),
      ),
    );
  }
}

class billerDash extends StatefulWidget {
  const billerDash({super.key});

  @override
  State<billerDash> createState() => _billerDashState();
}

class _billerDashState extends State<billerDash> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: ListView(
      children: [
        Center(
          child: Text("Biller"),
        ),
        Card(
          elevation: 5,
          child: Column(
            children: [
              Text(
                "Today Sales of ${email.replaceAll("@gmail.com", "").toUpperCase()}",
                style: TextStyle(fontSize: screenWidth * 0.04),
              ),
              RichText(
                  text: TextSpan(children: [
                TextSpan(
                    text: "Amount",
                    style: TextStyle(fontSize: screenWidth * 0.04)),
                TextSpan(
                    text: value, style: TextStyle(fontSize: screenWidth * 0.04))
              ]))
            ],
          ),
        )
      ],
    ));
  }
}

class managerDash extends StatefulWidget {
  const managerDash({super.key});

  @override
  State<managerDash> createState() => _managerDashState();
}

class _managerDashState extends State<managerDash> {
  @override
  void initState() {
    fetchInventryValue();
    // TODO: implement initState
    super.initState();
  }

  var TotalSPrice;
  stockValueA stockPrice = stockValueA();
  fetchInventryValue() async {
    TotalSPrice = await stockPrice.stockValueApi();
    print("Total Stock Price :${stockPrice.totalValue}");
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    screenWidth = MediaQuery.of(context).size.width;
    return Scaffold(
      body: Container(
        color: AppColors.bgColor,
        child: Center(
            child: Container(
          child: Card(
            elevation: 5,
            child: Container(
                // height: MediaQuery.of(context).size.height * 0.4,
                padding: EdgeInsets.all(20),
                decoration: BoxDecoration(
                  // color: AppColors.bgColor,
                  borderRadius: BorderRadius.all(Radius.circular(3)),
                ),
                child: Column(
                  children: [
                    Container(
                      margin: EdgeInsets.fromLTRB(0, 10, 0, 0),
                      child: Text(
                        "Welcome Manager :${email.replaceAll("@gamil.com", " ")}",
                        style: TextStyle(fontSize: screenWidth * 0.04),
                      ),
                    ),
                    Container(
                      margin: EdgeInsets.fromLTRB(0, 10, 0, 0),
                      child: Text(
                        "Total Stock Price :${stockPrice.totalValue}",
                        style: TextStyle(fontSize: 25),
                      ),
                    )
                  ],
                )),
          ),
        )),
      ),
    );
  }
}

class AdminnInventry extends StatefulWidget {
  const AdminnInventry({super.key});

  @override
  State<AdminnInventry> createState() => _AdminnInventryState();
}

class _AdminnInventryState extends State<AdminnInventry> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Card(
              elevation: 5,
              child: Container(
                padding: EdgeInsets.all(30),
                child: Text(
                  "Welcome to $userrole Page",
                  style: TextStyle(fontSize: screenWidth * 0.05),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
