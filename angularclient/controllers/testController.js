angular.module('myApp', ['testService']);

angular.module('myApp').controller('testController', ['$scope','testRequest',testController]);

function testController($scope, testRequest) {
	$scope.users={};
	$scope.user={};
	$scope.getAllUsers = function(){
		testRequest.users().success(function (data){
			$scope.users=data; // Asignaremos los datos de todos
			$scope.users.exist=1;
			$scope.user.exist=0;
		});
	}

	$scope.getUser = function(){
		testRequest.user($scope.user_id).success(function (data){
			$scope.user=data; // Asignaremos los datos
			$scope.user.exist=1;
			$scope.users.exist=1;
		});
	}

	$scope.selectId = function(id){
		$scope.user_id = id;
		$scope.getUser();
	}

	$scope.deleteUser = function(){
		testRequest.del($scope.user_id).success(function (data){
			alert("Usuario Eliminado.");
			$scope.user.exist=0;
			$scope.getAllUsers();

		});
	}

	$scope.editUser = function() {
		var firstname = $scope.user.firstname; //prompt("Enter the user firstname.");
		if(firstname == null){
			alert("El usuario debe tener Nombre.");
			return;
		}

		var lastname = $scope.user.lastname; //prompt("Enter the user lastname.");

		if(lastname == null){
			alert("El usuario debe tener Apellido.");

			return;
		}
		testRequest.edit($scope.user_id,firstname,lastname).success(function (){
			$scope.getAllUsers();
		});
	};

	$scope.add = function() {
		var firstname = $scope.user.firstname; //prompt("Enter the user firstname.");
		if(firstname == null){
			return;
		}
		var lastname = $scope.user.lastname; //prompt("Enter the user lastname.");
		if(lastname == null){
			return;
		}
		testRequest.add(firstname,lastname).success(function (){
			$scope.getAllUsers();
			//$(".alert").alert()
		});
	};

}
