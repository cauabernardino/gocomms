$('#signup-form').on('submit', createAccount);

function createAccount(event) {
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()) {
        alert("The passwords doesn't match!");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            username: $('#username').val(),
            password: $('#password').val(),
        }
    }).done(function () {
        alert("User has signed up!");
    }).fail(function (error) {
        console.log(error)
        alert("Error in signing up!");
    });
}