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

//This ajax/jquery code is being executed into the
//Teacher 1-4 aka Arithmos Dimosieuseon
ModipApp.arithmosDimosieuseon = function() {

  var originalValue;
  $('[contenteditable]').click(function(event) {
    originalValue = $(this).text();
  });

  $('[contenteditable]').focusout(function(event) {
    ajaxRequest($(this))
  });

  $('[contenteditable]').keydown(function(event) {
    //keycode 27 === esc
    if (event.which === 27) {
      //reset our field into its default value
      $(this).text(originalValue)
      $(this).trigger('blur');
    } else if (event.which === 13) {
      //this is the enter
      $(this).trigger('blur');
      originalValue = $(this).text()
      event.preventDefault();
      ajaxRequest($(this))
    }
  });

  //Due to some bug with the input we have to use this ugly hack
  //in order to use a select for the #teacherInput
  $('input').change(function(event) {
    var currentId = $(this).attr('id');
    if (currentId !== 'teacherInput') {
      return;
    }

    //ajax call here....
    ajaxRequest($(this))
  });

  //toggleUI is a virtual dumb class just for the selector
  $('.toggleUI').click(function(event) {
    var defaultType = $(this).text();
    var selectTypeElement = $(this).children('#selectType');
    var field6Element = $(this).children('#field6');

    //toggle the UI
    if (selectTypeElement.is(':hidden') && field6Element.is(':visible')) {
      selectTypeElement.show();
      field6Element.hide();
    } else {
      field6Element.show();
      selectTypeElement.hide();
    }

    selectTypeElement.on('change', function() {
      field6Element.text($(this).val())
      field6Element.show();
      //ajax call here
      ajaxRequest($(this))
    });

    selectTypeElement.focusout(function(event) {
      field6Element.show();
      selectTypeElement.hide();
    });
  });

  var ajaxRequest = function(_this) {
    var id = _this.data('id');
    var akademicYearId = _this.data('akademicyearid')
    var type = _this.data('type');
    var dataPostObject = {}

    var currentRow = _this.closest('tr');
    var tdElements = currentRow.find('td');
    $.each(tdElements, function(index, element) {
      var _self = $(this)
      var type =_self.data('type');
      if (type === undefined) {
        //this is a td element, so we must go into
        //the input element, so we must dig deeper.
        type = _self.children().data('type');

        //still empty? then it the select....
        if (type === undefined) {
          type = _self.children('select').data('type');
        }
      }

      switch (type) {
        case "author":
          dataPostObject.author = _self.text();
          break;
        case "title":
          dataPostObject.title = _self.text();
          break;
        case "is_magazine":
          dataPostObject.is_magazine = _self.children().is(':checked');
          break;
        case "publisher":
          dataPostObject.publisher = _self.text();
          break;
        case "publication_date":
          dataPostObject.publication_date = _self.children().val();
          break;
        case "type":
          dataPostObject.type = _self.children('select').val();
          break;
      }
    });


    $.post("/teacher/report/1/edit/" + id + "/" + akademicYearId, dataPostObject, function(data, status) {
      //do something with the callback
    })
  }

  $('.cloneRow.button').click(function() {
    //TODO For realtime cloning look below...
    //Just reload the page in order to copy the button
    //in the server side mode.
    location.reload();
  });
    /*$('.cloneRow.button').click(function(event) {
      //TODO This for the very brave people.....
      //Over 9000 work for the time being. Just do the "clone"
      //on server side mode.
      //http://stackoverflow.com/questions/20260222/jquery-remove-doesnt-work-on-clone
      //take the last row
      var lastRow = $('.arithmosDimosieuseonRow').last();
      //copy the last row
      var newRow = lastRow.clone();
      //append the cloned row into the end of the table
      newRow.insertAfter(lastRow)

      //Clear the cloned values from our previous row
      var tdElements = newRow.find('td');
      $.each(tdElements, function(index, element) {
        var _self = $(this)
        var type =_self.data('type');
        //increment the data-id by 1, otherwise data-id++
        var previousRowId = _self.data('id');
        if (previousRowId !== undefined) {
          var previousRowId = previousRowId + 1;
          _self.attr('data-id', previousRowId);
        }
        if (type === undefined) {
          //this is a td element, so we must go into
          //the input element, so we must dig deeper.
          type = _self.children().data('type');
          previousRowId = _self.children().data('id');
          previousRowId = previousRowId + 1;
          _self.children().attr('data-id', previousRowId);

          //still empty? then it the select....
          if (type === undefined) {
            type = _self.children('select').data('type');
            previousRowId =_self.children('select').data('id');
            previousRowId = previousRowId + 1;
            _self.children('select').attr('data-id', previousRowId);

          }
        }

        switch (type) {
          case "author":
            _self.text('');
            break;
          case "title":
            _self.text('');
            break;
          case "is_magazine":
            _self.children().prop('checked', false);
            break;
          case "publisher":
            _self.text('');
            break;
          case "publication_date":
            _self.children().val('');
            break;
          case "type":
            //DON'T REMOVE THIS.
            //there is no reason for our combobox to reset it to its
            //default value
            //_self.children('select').val('');
            break;
        }
      });
    });*/
}

ModipApp.ereunitikes_Ypodomes = function() {
  $('.button.cloneEntry').click(function(event) {

    //clone and append the new entry
    var trLastEntry =  $('.trEntry').last();
    var newTrEntry = trLastEntry.clone();
    newTrEntry.insertAfter(trLastEntry);

    var trEntries = $('.trEntry');
    var biggestIdHelper = -1;

    //find the biggerst Id.
    $.each(trEntries, function(index, element) {
      var _self =  $(this);
      var currentTextArea = _self.find('textarea');
      var currentNameAttr = currentTextArea.attr('name');
      var nameId = currentNameAttr.split('.')[1]
      if (nameId > biggestIdHelper) {
        biggestIdHelper = nameId
      }
    });

    var newId = parseInt(biggestIdHelper);
    newId += 1;
    var textArea = newTrEntry.find('textarea');
    //empty our field, since this is a new entry
    //remove the cloned value
    textArea.text('');
    //add a proper name to our new entry.
    textArea.attr('name','Helpers.' + newId + '.Content');
  });
}

ModipApp.ereunitikes_Ypodomes();
ModipApp.lessonMoveToDepartment();
ModipApp.akademicLesson();
ModipApp.arithmosDimosieuseon();
