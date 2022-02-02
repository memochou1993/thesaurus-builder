const title = document.querySelector('#title');
const root = document.querySelector('#root');
const subjectTemplate = document.querySelector('[data-subject-template]');
const descriptiveNoteTemplate = document.querySelector('[data-descriptive-note-template]');

/**
 * @param {HTMLElement} target
 * @param {Object} prop
 * @param {Object} prop.subject
 * @param {Object} prop.subject.term
 * @param {Array}  prop.subject.term.preferredTerms
 * @param {string} prop.subject.term.preferredTerms[].text
 * @param {Object} prop.subject.note
 * @param {Array}  prop.subject.note.descriptiveNotes
 * @param {string} prop.subject.note.descriptiveNotes[].text
 * @param {Array}  prop.children
 */
const render = (target, prop) => {
  setTimeout(() => {
    const [subject] = subjectTemplate.content.cloneNode(true).children;
    const [term, note, children] = subject.children;
    const [descriptiveNotes] = note.children;
    prop.subject.term.preferredTerms.forEach((item) => {
      term.textContent = item.text;
      term.classList.add(prop?.children?.length ? 'term-expandable' : 'term-expanded');
    });
    prop.subject.note?.descriptiveNotes?.forEach((item) => {
      const [descriptiveNote] = descriptiveNoteTemplate.content.cloneNode(true).children;
      const [text] = descriptiveNote.children;
      text.textContent = item.text;
      descriptiveNotes.append(descriptiveNote);
    });
    prop.children?.forEach((item) => render(children, item));
    target.appendChild(subject);
  }, 0);
};

(async () => {
  const data = await fetch('data.json').then((r) => r.json());
  title.textContent = data.title;
  render(root, data.root);
  document.body.hidden = false;
})();

root.addEventListener('click', (e) => {
  if (e.target.classList.contains('term-expandable')) {
    e.target.parentElement.querySelector('.children').classList.toggle('active');
    e.target.classList.toggle('term-expanded');
  }
});
