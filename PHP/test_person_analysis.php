<?php
use PHPUnit\Framework\TestCase;
require_once 'person_analysis.php';

class test_person_analysis extends TestCase {

    // Перевірка коректності розрахунку BMI
    public function testCalculateBMI() {
        $person = new person_analysis("John Doe", 25, 180, 75);
        $this->assertEquals(23.15, round($person->calculateBMI(), 2));
    }

    // Перевірка розрахунків станів здоров'я
    public function testGetHealthStatusNormalWeight() {
        $person = new person_analysis("John Doe", 25, 180, 75);
        $this->assertEquals("Normal weight", $person->getHealthStatus());
    }

    public function testGetHealthStatusUnderweight() {
        $person = new person_analysis("Jane Doe", 22, 160, 45);
        $this->assertEquals("Underweight", $person->getHealthStatus());
    }

    public function testGetHealthStatusOverweight() {
        $person = new person_analysis("Mark Smith", 30, 170, 80);
        $this->assertEquals("Overweight", $person->getHealthStatus());
    }

    public function testGetHealthStatusObesity() {
        $person = new person_analysis("Lucy Brown", 28, 160, 85);
        $this->assertEquals("Obesity", $person->getHealthStatus());
    }

    // Перевірка рядкового представлення об'єкта
    public function testToString() {
        $person = new person_analysis("John Doe", 25, 180, 75);
        $expectedString = "Name: John Doe, Age: 25, Height: 180, Weight: 75, BMI: 23.148148148148, Health Status: Normal weight";
        $this->assertEquals($expectedString, (string)$person);
    }



    // Хибний тест (зріст = 0)
    public function testCalculateBMIInvalidHeight() {
        $this->expectException(InvalidArgumentException::class);
        $person = new person_analysis("John Doe", 25, 0, 75);
        $person->calculateBMI();
    }

    // Хибний тест (вага < 0)
    public function testCalculateBMIInvalidWeight() {
        $this->expectException(InvalidArgumentException::class);
        $person = new person_analysis("John Doe", 25, 180, -75);
        $person->calculateBMI();
    }

    // Хибний тест (відсутнє ім'я)
    public function testMissingName() {
        $this->expectException(InvalidArgumentException::class);
        $person = new person_analysis("", 25, 180, 75);
        $person->__toString();
    }

    // Хибний тест (вік == null)
    public function testMissingAge() {
        $this->expectException(InvalidArgumentException::class);
        $person = new person_analysis("John Doe", null, 180, 75);
        $person->__toString();
    }

    // Хибний тест перевірка рядкового представлення об'єкта (дробове значення BMI округлене)
    public function testMissingToString() {
        $person = new person_analysis("John Doe", 25, 180, 75);
        $expectedString = "Name: John Doe, Age: 25, Height: 180, Weight: 75, BMI: 23.15, Health Status: Normal weight";
        $this->assertEquals($expectedString, (string)$person);
    }
}
?>