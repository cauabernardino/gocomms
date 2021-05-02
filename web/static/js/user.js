$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-profile').on('submit', editProfile);
$('#change-password').on('submit', changePassword)


function unfollow() {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error unfollowing user!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
        $('#unfollow').prop('disabled', false)
    })
}

function follow() {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userID}/follow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error following user!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
        $('#follow').prop('disabled', false)
    })

}


function editProfile(event) {
    event.preventDefault();

    $.ajax({
        url: "/edit-profile",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            username: $('#username').val(),
        }
    }).done(function () {
        Swal.fire({
            title: "Success!",
            text: "Profile edited successfully!",
            icon: "success",
            confirmButtonColor: "#4e4e50",
        }).then(function () {
            window.location = "/profile"
        });
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error editing profile!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
    })
}


function changePassword(event) {
    event.preventDefault();

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire({
            title: "Oops...",
            text: "The passwords don't match!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
        return;
    }

    $.ajax({
        url: "/change-password",
        method: "POST",
        data: {
            current: $('#current-password').val(),
            new: $('#new-password').val(),
        }
    }).done(function () {
        Swal.fire({
            title: "Success!",
            text: "Password was successfully changed!",
            icon: "success",
            confirmButtonColor: "#4e4e50",
        }).then(function () {
            window.location = "/profile"
        });
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error on changing password!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
    })
}