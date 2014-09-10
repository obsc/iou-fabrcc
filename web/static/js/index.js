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

	// Load transaction information
	var transactions, graphs;
	$.get("/transactions/json", function (data) { transactions = data; });
	console.log(transactions);
	$.get("/graph/json", function (data) { graphs=data; });
	console.log(graphs);

	setTimeout(showContent, 300);
}

function register() {

	$("#login").fadeTo(300, 0);

	// POST new user info
	user = {
		name : $("#username").val()
	};
	$.post("/users/new", user, function (data) {console.log(data);});

	// Load transaction information
	var transactions, graphs;
	$.get("/transactions/json", function (data) { transactions = data; });
	console.log(transactions);
	$.get("/graph/json", function (data) { graphs=data; });
	console.log(graphs);

	setTimeout(showContent, 300);
}

function showContent() {
	$('#login').toggleClass('hidden');
	$('#content').toggleClass('hidden');
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