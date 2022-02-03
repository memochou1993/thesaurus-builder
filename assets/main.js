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
 * @param {Object} prop.subject.terms
 * @param {Array}  prop.subject.terms[].text
 * @param {Array}  prop.subject.terms[].preferred
 * @param {Object} prop.subject.notes
 * @param {string} prop.subject.notes[].text
 * @param {Array}  prop.children
 */
const render = (target, prop) => {
  const [subject] = subjectTemplate.content.cloneNode(true).children;
  const [preferredTerm, notes, children] = subject.children;
  prop.subject.terms.forEach((item) => {
    if (item.preferred) {
      preferredTerm.textContent = item.text;
      preferredTerm.classList.add(prop?.children?.length ? 'preferred-term-expandable' : 'preferred-term-expanded');
    }
  });
  prop.subject.notes?.forEach((item) => {
    const [note] = noteTemplate.content.cloneNode(true).children;
    const [text] = note.children;
    text.textContent = item.text;
    notes.append(note);
  });
  target.appendChild(subject);
  setTimeout(() => prop?.children?.forEach((item) => render(children, item)), 0);
};

const toggleSpinner = async (delay = 0) => {
  await new Promise((res) => setTimeout(() => res(), delay));
  document.documentElement.classList.toggle('full-height');
  spinner.classList.toggle('hidden');
  app.classList.toggle('hidden');
}

(async () => {
  await toggleSpinner();
  const data = await fetch('data.json').then((r) => r.json());
  title.textContent = data.title;
  render(root, data.root);
  await toggleSpinner(1000);
})();

root.addEventListener('click', (e) => {
  if (e.target.classList.contains('preferred-term-expandable')) {
    e.target.parentElement.querySelector('.children').classList.toggle('hidden');
    e.target.classList.toggle('preferred-term-expanded');
  }
});
