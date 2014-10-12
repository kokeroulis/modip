describe('Teacher', function () {

  describe('Asset List', function() {

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
      var expected = {
        teacher: {
          id: 1,
          name: 'superteacher1',
          email: 'superteacher1@teilar.gr',
          department: {
            id: 1,
            name: 'T.P.T.'
          }
        }, assetTypes: [
          {
            id: 1,
            name: 'Γενικά',
            assets: []
          },
          {
            id: 2,
            name: 'Βιβλία',
            assets: [
              {
                id: 1,
                content: 'Awesome book',
                teacher: {
                  id: 0,
                  name: '',
                  email: '',
                  department: { id: 0, name: '' }
                },
                assetType: 2 }
            ]
          },
          {
            id: 3,
            name: 'Δημοσιεύσεις',
            assets: []
          }
        ]
      };

      service.info().then(function(data) {
        expect(data).toEqual(expected)
        done();
      });
    });
  });

});

