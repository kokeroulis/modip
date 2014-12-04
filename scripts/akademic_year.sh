#!/usr/bin/env bash

add_lesson_akademic_year() {
    file_number=$1
    origin_number=$2
    origin_dollar='$'
    origin_dollar+=$origin_number
    new_number=$((origin_number + 1))
    new_dollar='$'
    new_dollar+=$new_number

    echo $file_number
    echo $origin_number
    echo $origin_dollar
    echo $new_dollar

    sed -i "s,WHERE lesson = ${origin_dollar}, WHERE lesson = ${origin_dollar} AND akademic_year = ${new_dollar}," models/forms/lesson_entry${file_number}.go
    sed -i 's,WHERE lesson = $1 , WHERE lesson = $1 AND akademic_year = $2,' models/forms/lesson_entry${file_number}.go
    sed -i 's/Load(lessonId int)/Load(lessonId int, akademicYearId int)/' models/forms/lesson_entry${file_number}.go
    sed -i 's/Update(lessonId int)/Update(lessonId int, akademicYearId int)/' models/forms/lesson_entry${file_number}.go
    sed -i 's/lessonId)/lessonId,\n\takademicYearId)/' models/forms/lesson_entry${file_number}.go
    sed -i 's/lesson [ ]* int references lesson(id) on delete cascade,/lesson int references lesson(id) on delete cascade,\n\takademic_year int references akademic_year(id) on delete cascade,/' db/schema/lesson_create_entry${file_number}.go

    gofmt -w=true models/forms/lesson_entry${file_number}.go
}

add_teacher_akademic_year() {
    file_number=$1
    origin_number=$2
    origin_dollar='$'
    origin_dollar+=$origin_number
    new_number=$((origin_number + 1))
    new_dollar='$'
    new_dollar+=$new_number

    echo $file_number
    echo $origin_number
    echo $origin_dollar
    echo $new_dollar

    sed -i "s,WHERE teacher = ${origin_dollar}, WHERE teacher = ${origin_dollar} AND akademic_year = ${new_dollar}," models/forms/teacher_entry${file_number}.go
    sed -i 's,WHERE teacher = $1 , WHERE teacher = $1 AND akademic_year = $2,' models/forms/teacher_entry${file_number}.go
    sed -i 's/Load(teacherId int)/Load(teacherId int, akademicYearId int)/' models/forms/teacher_entry${file_number}.go
    sed -i 's/Update(teacherId int)/Update(teacherId int, akademicYearId int)/' models/forms/teacher_entry${file_number}.go
    sed -i 's/teacherId)/teacherId,\n\takademicYearId)/' models/forms/teacher_entry${file_number}.go
    sed -i 's/teacher [ ]* int references teacher(id) on delete cascade,/teacher int references teacher(id) on delete cascade,\n\takademic_year int references akademic_year(id) on delete cascade,/' db/schema/teacher_create_entry${file_number}.go

    gofmt -w=true models/forms/teacher_entry${file_number}.go
}


add_lesson_akademic_year 1 3
add_lesson_akademic_year 2 20
add_lesson_akademic_year 3 6
add_lesson_akademic_year 4 3
add_lesson_akademic_year 5 7

add_lesson_akademic_year 11 6
add_lesson_akademic_year 12 4
add_lesson_akademic_year 13 2
add_lesson_akademic_year 14 8
add_lesson_akademic_year 15 3
add_lesson_akademic_year 16 2

add_teacher_akademic_year 2 11
add_teacher_akademic_year 3 5
add_teacher_akademic_year 5 2
