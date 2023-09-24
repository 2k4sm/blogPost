document.addEventListener('DOMContentLoaded', function() {
    const markdownInput = document.querySelector('#markdown-input');
    const markdownPreview = document.querySelector('#markdown-preview');
  
    markdownInput.addEventListener('input', updatePreview);
  
    function updatePreview() {
      const markdownText = markdownInput.value;
      markdownPreview.innerHTML = marked.parse(markdownText);
    }
});