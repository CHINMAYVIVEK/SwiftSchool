class Data {
  final int id;
  final String value;

  Data({required this.id, required this.value});

  factory Data.fromJson(Map<String, dynamic> json) {
    return Data(
      id: json['id'],
      value: json['value'],
    );
  }
}
