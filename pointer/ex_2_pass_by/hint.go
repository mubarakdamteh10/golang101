package main

import "fmt"

func incrementByValueHint(num int) {
	num = num + 1
	// ฟังกืชันนี้คือการ pass by value
	// ซึ่งจะมอง num เป็นตัวแปรใหม่และคำนวณเพียงแค่ในฟังก์ชันนี้แล้วจบไป
	// ไม่มีการเปลี่ยนแปลงค่าของตัวแปรที่ใส่ใน parameter แต่อย่างใด
}

func incrementByReferenceHint(num *int) {
	*num = *num + 1
	// เข้าถึงค่าผ่าน pointer และเพิ่มค่าไปอีก 1
	// ตัวแปรที่ส่งใน parameter จะถูกเปลี่ยนค่าไปด้วย
	// เรียกว่า pass by reference
}

func hint() {

	x := 10
	fmt.Println("ค่าเริ่มต้นของ x:", x)

	incrementByValueHint(x)                          // เรียกใช้ฟังก์ชันแบบ Pass-by-Value
	fmt.Println("หลังจากเรียก incrementByValue:", x) // คาดหวังว่า x ยังควรเป็น 10 (ไม่เปลี่ยนแปลง)

	incrementByReferenceHint(&x)                       // เรียกใช้ฟังก์ชันแบบ Pass-by-Reference โดยส่งที่อยู่หน่วยความจำของ x
	fmt.Println("หลังจากเรียก incrementByPointer:", x) // คาดหวังว่า x ควรเปลี่ยนเป็น 11
}
