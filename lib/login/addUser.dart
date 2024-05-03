import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:medapp/apimethods/addUserM.dart';

import '../colors.dart';
import '../method/stockviewM.dart';

class UserAddView extends StatefulWidget {
  const UserAddView({super.key});

  @override
  State<UserAddView> createState() => _UserAddViewState();
}

class _UserAddViewState extends State<UserAddView> {
  String? Role;
  AdduserA useradd = AdduserA();
  @override
  void initState() {
    // TODO: implement initState
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;
    TextEditingController _username = TextEditingController();
    TextEditingController _password = TextEditingController();

    ScaffoldMsg(String message) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          backgroundColor: Color.fromARGB(255, 43, 41, 41),
          content: RichText(
              text: TextSpan(children: [
            TextSpan(
              text: message,
              style: TextStyle(
                  fontSize: screenWidth * 0.03,
                  color: Color.fromARGB(255, 255, 255, 255)),
            ),
          ])),
          duration: Duration(seconds: 3),
        ),
      );
    }

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Container(
              color: AppColors.bgColor,
              width: screenWidth * 0.6,
              child: Column(
                children: [
                  Container(
                    margin: EdgeInsets.fromLTRB(0, 25, 0, 0),
                    width: screenWidth * 0.5,
                    child: DropdownButtonFormField<String>(
                      value: Role,
                      onChanged: (String? newValue) {
                        setState(() {
                          Role = newValue!;
                        });
                      },
                      items: [
                        DropdownMenuItem<String>(
                          value: "admin",
                          child: Text("Admin"),
                        ),
                        DropdownMenuItem<String>(
                          value: "manager",
                          child: Text("Manager"),
                        ),
                        DropdownMenuItem<String>(
                          value: "inventry",
                          child: Text("Inventry"),
                        ),
                        DropdownMenuItem<String>(
                          value: "biller",
                          child: Text("Biller"),
                        )
                      ],
                    ),
                  ),
                  Container(
                    width: screenWidth * 0.5,
                    child: TextField(
                      controller: _username,
                      decoration: InputDecoration(
                          hintText: "User Name as Email",
                          border: UnderlineInputBorder(
                            borderSide: BorderSide(
                              color: Colors.black,
                              width: 1.0,
                            ),
                          )),
                    ),
                  ),
                  Container(
                    width: screenWidth * 0.5,
                    child: TextField(
                      controller: _password,
                      decoration: InputDecoration(
                          hintText: "Password",
                          border: UnderlineInputBorder(
                            borderSide: BorderSide(
                              color: Colors.black,
                              width: 1.0,
                            ),
                          )),
                    ),
                  ),
                  Row(
                    children: [
                      MaterialButton(
                        onPressed: () {
                          var user = _username.text;
                          var pass = _password.text;
                          if (Role == null) {
                            ScaffoldMsg("Please Select Role");
                          } else if (!user.contains("@") &&
                              !user.contains(".com")) {
                            ScaffoldMsg("Please Enter valid username");
                          } else if (_password.text.isEmpty) {
                            ScaffoldMsg("Please fill Password");
                          } else if (_username.text.isEmpty) {
                            ScaffoldMsg("Please insert username");
                          } else if (pass.length < 8) {
                            ScaffoldMsg(
                                "Password must be at least 8 characters");
                          } else {
                            print("insert");
                            Role = null;

                            _username.clear();
                            _password.clear();
                            setState(() {});
                          }
                        },
                        child: MaterialButton(
                          onPressed: () {
                            Future<dynamic> resp = useradd.AdduserAPI(
                                _username.text, _password.text, Role!);
                            print(resp.toString());
                            if (resp != null) {
                              ScaffoldMsg() {
                                ScaffoldMessenger.of(context).showSnackBar(
                                  SnackBar(
                                    backgroundColor:
                                        Color.fromARGB(255, 92, 228, 99),
                                    content: RichText(
                                        text: TextSpan(children: [
                                      TextSpan(
                                        text: "User Added",
                                        style: TextStyle(
                                            fontSize: screenWidth * 0.03,
                                            color:
                                                Color.fromARGB(255, 0, 0, 0)),
                                      ),
                                    ])),
                                    duration: Duration(seconds: 3),
                                  ),
                                );
                              }
                            }
                          },
                          child: Container(
                              margin: EdgeInsets.fromLTRB(0, 20, 0, 20),
                              padding: EdgeInsets.fromLTRB(10, 5, 10, 5),
                              decoration: BoxDecoration(
                                color: AppColors.buttonColor,
                                borderRadius: BorderRadius.circular(20),
                              ),
                              child: Text("Add User")),
                        ),
                      ),
                    ],
                  )
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
