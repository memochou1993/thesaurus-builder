const app = document.querySelector('#app');
const spinner = document.querySelector('#spinner');
const title = document.querySelector('#title');
const root = document.querySelector('#root');
const subjectTemplate = document.querySelector('[data-subject-template]');
const noteTemplate = document.querySelector('[data-note-template]');

/**
 * @param {HTMLElement} target
 * @param {Object} prop
 * @param {Object} prop.subject
 * @param {Object} prop.subject.term
 * @param {Array}  prop.subject.term.preferredTerms
 * @param {string} prop.subject.term.preferredTerms[].text
 * @param {Object} prop.subject.notes
 * @param {string} prop.subject.notes[].text
 * @param {Array}  prop.children
 */
const render = (target, prop) => {
  setTimeout(() => {
    const [subject] = subjectTemplate.content.cloneNode(true).children;
    const [term, notes, children] = subject.children;
    prop.subject.term.preferredTerms.forEach((item) => {
      term.textContent = item.text;
      term.classList.add(prop?.children?.length ? 'term-expandable' : 'term-expanded');
    });
    prop.subject.notes?.forEach((item) => {
      const [note] = noteTemplate.content.cloneNode(true).children;
      const [text] = note.children;
      text.textContent = item.text;
      notes.append(note);
    });
    prop.children?.forEach((item) => render(children, item));
    target.appendChild(subject);
  }, 0);
};

(async () => {
  const data = await fetch('data.json').then((r) => r.json());
  title.textContent = data.title;
  render(root, data.root);
  await new Promise((res) => setTimeout(() => res(), 1000));
  spinner.classList.toggle('hidden');
  document.documentElement.classList.remove('full-height');
  app.classList.toggle('hidden');
})();

root.addEventListener('click', (e) => {
  if (e.target.classList.contains('term-expandable')) {
    e.target.parentElement.querySelector('.children').classList.toggle('hidden');
    e.target.classList.toggle('term-expanded');
  }
});
