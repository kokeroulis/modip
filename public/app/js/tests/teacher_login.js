describe('Teacher', function () {

  describe('Login', function() {

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
        id:1,
        name:'superteacher1',
        email:'superteacher1@teilar.gr',
        department: {
          id:1,
          name: 'T.P.T.'
        }
      };

      service.login('superteacher1@teilar.gr', 'superteacher1').then(function(data) {
        expect(data).toEqual(expected)
        done();
      });
    });

    it('should fail', function(done) {
      var expected = {
        body: {
          Code: 401,
          Name: 'Authorization Failed'
        }
      };

      service.login('superteacher1@teilar.gr', 'superteacher12').then(function(data) {},
      function(error) {
        expect(error).toEqual(expected)
        done();
      });
    });
  });

});

