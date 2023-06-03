const postsDiv = document.querySelector('.posts');
const signupBtn = document.querySelector('#signupBtn');
const loginBtn = document.querySelector('#loginBtn');
const logoutBtn = document.querySelector('#logoutBtn');
const createPostBtn = document.querySelector('#createPostBtn');
const categoryBtn = document.querySelector('#categoryBtn');
const categoryList = document.createElement('ul');
categoryList.classList.add('category-list');
let isCategoryListVisible = false;
  
signupBtn.addEventListener('click', () => {
    window.location.href = 'http://localhost:8080/signup';
});

loginBtn.addEventListener('click', () => {
    window.location.href = 'http://localhost:8080/login';
});
 
logoutBtn.addEventListener('click', () => {
    window.location.href = 'http://localhost:4000/logout';
});

createPostBtn.addEventListener('click', () => {
    window.location.href = 'http://localhost:8080/post';
});

categoryBtn.addEventListener('click', () => {
    if (!isCategoryListVisible) {
        categoryList.innerHTML = '';
        const categories = ["Anime", "Cars", "Video Games", "Programming", "Sport", "Politics", "Algorithm", "Problems"];
        categories.forEach(category => {
            const listItem = document.createElement('li');
            listItem.textContent = category;
            listItem.addEventListener('click', () => {
                const chosenCategory = encodeURIComponent(category);
                window.location.href = `http://localhost:8080/post/category?category=${chosenCategory}`;
            });
            categoryList.appendChild(listItem);
        });
        categoryBtn.appendChild(categoryList);
    } else {
        categoryList.innerHTML = '';
    }
    isCategoryListVisible = !isCategoryListVisible;
});

// Make GET request to retrieve posts
fetch('http://localhost:4000/home')
    .then(response => response.json())
    .then(posts => {
        // Displaying the posts
        posts.forEach(post => {
            const postElement = document.createElement('div');
            postElement.classList.add('post');
            postElement.innerHTML = `
                <h2>${post.title}</h2>
                <p>Categories: ${post.category.join(', ')}</p>
                <p>Author: ${post.user_name}</p>
            `;

            postElement.addEventListener('click', () => {
                const postId = encodeURIComponent(post.id);
                window.location.href = `http://localhost:8080/post?id=${postId}`;
            });

            postsDiv.appendChild(postElement);
        });
    })
    .catch(error => {
        console.log('Error retrieving posts:', error);
    });

    document.addEventListener("DOMContentLoaded", function() {
      console.log('waaaaa');
      const loginButton = document.getElementById("signupBtn");
      const signupButton = document.getElementById("loginBtn");
      const logoutButton = document.getElementById("logoutBtn");
    
      // Function to check if the cookie exists
      function checkCookieExistence(cookieName) {
        const cookies = document.cookie.split(";").map(cookie => cookie.trim());
        console.log(cookies.includes(cookieName))
        return cookies.includes(cookieName);
      }
    
      if (checkCookieExistence("Session")) {
        loginButton.style.display = "none";
        signupButton.style.display = "none";
        logoutButton.style.display = "block";
      } else {
        loginButton.style.display = "block";
        signupButton.style.display = "block";
        logoutButton.style.display = "none";
      }
    });