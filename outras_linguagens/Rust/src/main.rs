pub mod tempconv;

use tempconv::Temperatures;
fn main() {
    println!("Que frio! {:?}", Temperatures::Celsius(0.0).absolute_zero());
    println!("Fervendo! {:?}", Temperatures::Celsius(0.0).to_fahrenheit().boiling());
}
