fn main() {
    let mut message = "Hello".to_string(); // Обратите внимание на `mut`
    println!("До изменения: {}", message);
    change_message(&mut message); // Передаём изменяемую ссылку
    println!("После изменения: {}", message);

    // неизменяемая ссылка создается и используется ДО создания изменяемой ссылки.
    let mut s1 = "hello".to_string();
    let s3 = &s1;   // Неизменяемая ссылка (посмотрели на квартиру)
    println!("{}", s3); // hello
    // Дальше неизменяемая ссылка s3 НЕ используется
    let s2 = &mut s1;   // Изменяемая ссылка (отдали ключи ремонтникам)
    s2.push('!');
    println!("{}", s1); // hello!
    
}

fn change_message(mes: &mut String) {
    mes.push('!'); // Теперь ошибки нет
}
