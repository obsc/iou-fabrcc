var userMap, graph, history;
var rowsShown = 0;

$.get("/users/json", null, function (data) { mapSetup(JSON.parse(data)); });
$.get("/transactions/json", null, function (data) { transactions = JSON.parse(data); });
$.get("/graph/json", null, function (data) { graph = JSON.parse(data); });

function listTransactions(transactions) {

	// one placeholder row for a new one
	var add = document.createElement('div');
	add.className = 't_row';
	$("#transactions").appendChild(add);

	// iterate over all transactions and list them
	var i;
	for (i in transactions) {
		var add = document.createElement('div');
		add.className = 't_row';
		$("#transactions").appendChild(add);
	}
}

function mapSetup(users) {
	console.log(users);
	userMap = {};
	users.forEach( function(user) {
		userMap[user.Name] = user.UserId;
	});

	// Setup autocomplete
	var names = Object.keys(userMap);
	$('#source').autocomplete({
		source: names
	});
	$('#sink').autocomplete({
		source: names
	});
}

function register() {

	// POST new user info
	user = {
		name : $("#username").val()
	};

	$.post("/users/new", user, function (data) {console.log(data);});

	$.get("/users/json", null, function (data) { mapSetup(JSON.parse(data)); });
	$.get("/transactions/json", null, function (data) { transactions = JSON.parse(data); });
	$.get("/graph/json", null, function (data) { graph = JSON.parse(data); });

}

function expand() {
	rowsShown += 5;
	var height = $('#transactions').height();
	$('#transactions').height(height + 50);
}

function showContent() {
	$('#user').toggleClass('hidden');
	$('#content').toggleClass('hidden');
	
	// Content-related changes
	$('#expand').click(expand);
}

function newTransaction() {

	transaction = {
		source : userMap[$('#source').val()],
		sink : userMap[$('#sink').val()],
		value : (+($('#amount').val())) * 100,
		reason : $('#reason').val()
	};
	$.post("/transactions/new", transaction);

}