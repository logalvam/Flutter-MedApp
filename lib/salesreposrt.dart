import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:medapp/apimethods/salesReportM.dart';
import 'package:medapp/colors.dart';
import 'package:medapp/login/loginPage.dart';

import 'method/salesReportM.dart';

class SalesReport extends StatefulWidget {
  const SalesReport({super.key});

  @override
  State<SalesReport> createState() => _SalesReportState();
}

SalesReportA reportA = SalesReportA();

class _SalesReportState extends State<SalesReport> {
  Salesreport? reportlist;
  DateTime fromDate = DateTime.now();
  DateTime toDate = DateTime.now();

  var from;
  var to;

  _selectDate(BuildContext context) async {
    final DateTime? picked = await showDatePicker(
      context: context,
      initialDate: fromDate,
      firstDate: DateTime(2000),
      lastDate: DateTime.now(),
    );
    if (picked != null && picked != toDate)
      setState(() {
        fromDate = picked;
      });
  }

  _todate(BuildContext context) async {
    final DateTime? picked1 = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: fromDate,
      lastDate: DateTime.now(),
    );

    if (picked1 != null && picked1 != toDate)
      setState(() {
        toDate = picked1;
        from = fromDate.toString().replaceAll("00:00:00.000", "");
        to = toDate.toString().replaceAll("00:00:00.000", "");
        print("$from , $to");
      });
  }

  @override
  void initState() {
    reportlist = Salesreport(salesarr: []);
    // TODO: implement initState
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    TextStyle fonts = TextStyle(
      fontSize: screenWidth * 0.02,
      fontFamily: "Helvetica",
    );
    return Scaffold(
      body: ListView(
        children: [
          Column(
            // mainAxisSize: MainAxisSize.min,
            children: <Widget>[
              Row(
                children: [
                  Text(
                    "${fromDate.toLocal()}".split(' ')[0],
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                  SizedBox(width: 10),
                  Text(
                    "${toDate.toLocal()}".split(' ')[0],
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                ],
              ),
              // Divider(),
              Row(
                children: [
                  MaterialButton(
                    onPressed: () => _selectDate(context), // Refer step 3
                    child: Container(
                      padding: EdgeInsets.all(10),
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(20),
                        color: AppColors.buttonColor,
                      ),
                      child: Text(
                        'From date',
                        // style: TextStyle(
                        //     color: Colors.black, fontWeight: FontWeight.bold),
                      ),
                    ),
                  ),
                  MaterialButton(
                    onPressed: () => _todate(context), // Refer step 3
                    child: Container(
                      padding: EdgeInsets.all(10),
                      decoration: BoxDecoration(
                        color: AppColors.buttonColor,
                        borderRadius: BorderRadius.circular(20),
                      ),
                      child: Text(
                        'To date',
                        // style: TextStyle(),
                      ),
                    ),
                    // color: AppColors.buttonColor,
                  ),
                ],
              ),
              MaterialButton(
                onPressed: () async {
                  var res = await reportA.SalesReportApi(from, to);
                  setState(() {
                    reportlist = res;
                    // reportlist.salesarr[0]
                  });
                },
                child: Container(
                  padding: EdgeInsets.all(10),
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(20),
                      color: AppColors.buttonColor),
                  child: Text("Search"),
                ),
              ),
              Container(
                width: screenWidth,
                child: DataTable(
                    columnSpacing: 4,
                    columns: <DataColumn>[
                      DataColumn(
                          tooltip: "Bill no",
                          label: Text(
                            "Bill No",
                            style: TextStyle(
                                fontSize: screenWidth * 0.03,
                                fontStyle: FontStyle.italic,
                                fontWeight: FontWeight.bold),
                          )),
                      DataColumn(
                          label: Text(
                        "Bill Date",
                        style: TextStyle(
                            fontSize: screenWidth * 0.03,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Medicine name",
                        style: TextStyle(
                            fontStyle: FontStyle.italic,
                            fontSize: screenWidth * 0.03,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Quantity",
                        style: TextStyle(
                            fontSize: screenWidth * 0.03,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                      DataColumn(
                          label: Text(
                        "Amount",
                        style: TextStyle(
                            fontSize: screenWidth * 0.03,
                            fontStyle: FontStyle.italic,
                            fontWeight: FontWeight.bold),
                      )),
                    ],
                    rows: List<DataRow>.generate(
                        reportlist!.salesarr.length,
                        (index) => DataRow(cells: <DataCell>[
                              DataCell(
                                Text(
                                  reportlist!.salesarr[index].billno.toString(),
                                  style: fonts,
                                ),
                              ),
                              DataCell(
                                Text(
                                  reportlist!.salesarr[index].billdate
                                      .toString()
                                      .replaceAll("00:00:00.000", ""),
                                  style: fonts,
                                ),
                              ),
                              DataCell(Text(reportlist!.salesarr[index].medname,
                                  style: fonts)),
                              DataCell(Text(
                                (reportlist!.salesarr[index].quantity)
                                    .toString(),
                                style: fonts,
                              )),
                              DataCell(Text(
                                (reportlist!.salesarr[index].amount).toString(),
                                style: fonts,
                              )),
                            ]))),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
