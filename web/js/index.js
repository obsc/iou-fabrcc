var username;
var password;

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

function login() {

	$("#login").fadeTo(300, 0);

	// Identify the user from fields
	username = $("#username").val();
	password = $("#password").val();

	// Load transaction information
	var transactions, graphs;
	$.get("/transactions/json", username, function (data) { listTransactions(data); });
	$.get("/graphs/json", username, function (data) { graphs=data; });

	// Reveal transaction information
}

function register() {

	$("#login").fadeTo(300, 0);

	// Identify the user from fields
	username = $("#username").val();
	password = $("#password").val();

	// POST new user info
	$.post("/users/new", username, function (data) {console.log(data);});
}

function newTransaction() {

	$.post("transactions/new", function (data) { listTransactions(data); });

}