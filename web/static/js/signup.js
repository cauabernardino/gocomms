$('#signup-form').on('submit', createAccount);

function createAccount(event) {
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire({
            title: "Oops...",
            text: "The passwords don't match!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        })
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
        Swal.fire({
            title: "Welcome!",
            text: "You signed up successfully!",
            icon: "success",
            confirmButtonColor: "#4e4e50",
        }).then(function () {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val(),
                }

            }).done(function () {
                window.location = "/home";
            }).fail(function () {
                Swal.fire({
                    title: "Oops...",
                    text: "Error authenticating the user!",
                    icon: "error",
                    confirmButtonColor: "#4e4e50",
                })
            })
        });
    }).fail(function (error) {
        Swal.fire({
            title: "Oops...",
            text: "Error in signing up!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        })
    });
}