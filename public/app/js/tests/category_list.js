describe('Teacher', function () {

  describe('Category List', function() {

    var tester, service;

    beforeEach(function() {
      tester = ngMidwayTester('modipApp');
      service = tester.inject('TeacherService');
    });

    afterEach(function() {
      tester.destroy();
      tester = null;
    });

    it('should succeed', function(done) {
      var expected = [{
        id: 1,
        name: "Διδάσκων",
        categories: [
          {
            id: 1,
            name: "Α.Δ. Εκπ. Προσωπικού",
            children: [
              {
                id: 1001,
                name: "Αριθμός Δημοσιεύσεων",
                children: null,
                authActions: {
                  is_visible_to_teacher: false,
                  teacher_can_edit: false
                }
              },
              {
                id: 1002,
                name: "Αναγνώριση του επιστημονικού και άλλου έργου",
                children: null,
                authActions: {
                  is_visible_to_teacher: false,
                  teacher_can_edit: false
                }
              },
              {
                id: 1003,
                name: "Ερευνητικά προγράμματα και έργα",
                children: null,
                authActions: {
                  is_visible_to_teacher: false,
                  teacher_can_edit: false
                }
              },
              {
                id: 1004,
                name: "Ερευνητικές Υποδομές",
                children: null,
                authActions: {
                  is_visible_to_teacher: false,
                  teacher_can_edit: false
                }
              },
              {
                id: 1005,
                name: "Σύνδεση με την κοινωνία",
                children: null,
                authActions: {
                  is_visible_to_teacher: false,
                  teacher_can_edit: false
                }
              }
            ],
            authActions: {
              is_visible_to_teacher: false,
              teacher_can_edit: false
            }
        }
        ]
      }];

      service.categoryList().then(function(data) {
        expect(data).toEqual(expected)
        done();
      });
    });
  });

});

