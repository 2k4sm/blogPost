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
          const img = mutation.target.querySelector('img');
          if (img) {
              img.style.height = '10vh';
          }
      });
  });

  observer.observe(document.body, { childList: true, subtree: true });
});
