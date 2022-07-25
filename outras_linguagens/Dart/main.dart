import 'tempconv/tempconv.dart';

void main() {
  final Tempconv tempconv = Tempconv();
  print('Que frio! ${Tempconv.absoluteZeroC}');
  print('Fervendo! ${tempconv.CtoF(Tempconv.boilingC)}');
}
