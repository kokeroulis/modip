ModipApp.lessonMoveToDepartment = function() {
  $('#lessonMoveToDepartment').change(function() {
    window.location.href = $(this).val()
  })

}

ModipApp.akademicLesson = function() {

  var lessonId = null;
  var departmentId = null;
  $('#lessonAkademic').change(function() {
    lessonId = $(this).val();

    if (lessonId && departmentId) {
      var url = "/lesson/list/pre/degree/" + String(departmentId) + "/" + String(lessonId)
      window.location.href = url
    }
  })

  $('#departmentAkademic').change(function() {
    departmentId = $(this).val();

    if (lessonId && departmentId) {
      var url = "/lesson/list/pre/degree/" + String(departmentId) + "/" + String(lessonId)
      window.location.href = url
    }
  })
}

ModipApp.lessonMoveToDepartment();
ModipApp.akademicLesson();
