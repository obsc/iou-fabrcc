var users, graph, history;
var rowsShown = 0;

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

// Login Menu functions
function login() {

	$("#login").fadeTo(300, 0);

	$.get("/users/json", null, function (data) { users = JSON.parse(data); });
	$.get("/transactions/json", null, function (data) { transactions = JSON.parse(data); });
	$.get("/graph/json", null, function (data) { graph = JSON.parse(data); });

	setTimeout(showContent, 300);
}

function register() {

	$("#login").fadeTo(300, 0);

	// POST new user info
	user = {
		name : $("#username").val()
	};
	$.post("/users/new", user, function (data) {console.log(data);});

	$.get("/users/json", null, function (data) { users = JSON.parse(data); });
	$.get("/transactions/json", null, function (data) { transactions = JSON.parse(data); });
	$.get("/graph/json", null, function (data) { graph = JSON.parse(data); });

	setTimeout(showContent, 300);
}

function expand() {
	rowsShown += 5;
	var height = $('#transactions').height();
	$('#transactions').height(height + 50);
}

function showContent() {
	$('#user').toggleClass('hidden');
	$('#content').toggleClass('hidden');
	$('#expand').click(expand);
}

function newTransaction() {

	transaction = {
		source : sourceId,
		sink : sinkId,
		value : value,
		reason : reason
	};

	$.post("transactions/new", function (data) { listTransactions(data); });

}