describe('Teacher', function () {

  describe('Asset Move', function() {

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
        id: 1,
        content: 'Super Book',
        teacher: {
          id: 1,
          name: 'superteacher1',
          email: 'superteacher1@teilar.gr',
          department: {
            id: 1,
            name: 'T.P.T.'
          }
        },
        assetType: 2
      };

      service.moveAsset(1, 2).then(function(data) {
        console.log(data)
        expect(data).toEqual(expected)
        done();
      });
    });

    it('should fail', function(done) {
      var expected = {
        body: {
          Code: 200,
          Name: 'InvalidAsset'
        }
      };

      service.moveAsset(1000, 2).then(function(data) {},
      function(error) {
        expect(error).toEqual(expected)
        done();
      });
    });
  });

});

