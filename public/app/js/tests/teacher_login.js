describe('Teacher', function () {

  describe('Login', function() {

    var tester, service;

    beforeEach(function(){
      tester = ngMidwayTester('modipApp');
      service = tester.inject('TeacherService');
    });

    afterEach(function(){
      tester.destroy();
      tester = null;
    });

    it('should make a real XHR GET for http-hello.html', function(done) {
      var expected = {
        id:1,
        name:"superteacher1",
        email:"superteacher1@teilar.gr",
        department: {
          id:1,
          name: "T.P.T."
        }
      };

      service.login('superteacher1@teilar.gr', 'superteacher1').then(function(data) {
        expect(data).toEqual(expected)
        done();
      })
    });

  });

});

