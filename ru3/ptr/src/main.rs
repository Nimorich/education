fn main() {
    let mut s1 = "hello".to_string();
    {
        let s3 = &mut s1; // Отдали ключи первой бригаде
        s3.push('?');
    } // Первая бригада закончила и ушла
    let s2 = &mut s1;  // Отдали ключи второй бригаде
    s2.push('!');
    println!("{s1}");  // hello?!
}
