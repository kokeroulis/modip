ModipApp.lessonMoveToDepartment = function() {
  $('#lessonMoveToDepartment').change(function() {
    window.location.href = $(this).val()
  })

}

ModipApp.lessonMoveToDepartment();
