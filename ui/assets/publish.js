document.addEventListener('DOMContentLoaded', function() {
    const markdownInput = document.querySelector('#markdown-input');
    const markdownPreview = document.querySelector('#markdown-preview');
  
    markdownInput.addEventListener('input', updatePreview);
  
    function updatePreview() {
      const markdownText = markdownInput.value;
      markdownPreview.innerHTML = marked.parse(markdownText);
    }
});

document.addEventListener('DOMContentLoaded', function() {
  const observer = new MutationObserver(function(mutations) {
      mutations.forEach(function(mutation) {
          const images = mutation.target.querySelectorAll('img');
          images.forEach(function(img) {
              img.style.height = '20vh';
              img.style.width = 'auto';
          });
      });
  });

  observer.observe(document.body, { childList: true, subtree: true });
});
