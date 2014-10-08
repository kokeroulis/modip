describe('When calling SUT', function () {

  var tester, service;

  describe('with midway tester', function() {

    beforeEach(function(){
      tester = ngMidwayTester('modipApp');
      service = tester.inject('TeacherService');
    });

    afterEach(function(){
      tester.destroy();
      tester = null;
    });

    it('should make a real XHR GET for http-hello.html', function() {
      var expected = "Hello, $http!";
      console.log("real GET http-hello.html test");
      console.log(service.login('sadsdads', 'dsasadsad'))
    });
  });
});

