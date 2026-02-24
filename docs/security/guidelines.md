# Security Guidelines for Client Projects

## ⚠️ CRITICAL: What NOT to Push to GitHub

### 🔴 Never Push These Files

**1. Environment Files & Secrets**

- ❌ `.env` (production credentials)
- ❌ `.env.local`
- ❌ `.env.production`
- ❌ Any file containing passwords, API keys, JWT secrets
- ❌ Service account keys (`.json`, `.pem`, `.key`)

**2. Database Files**

- ❌ `*.db`, `*.sqlite`, `*.sqlite3`
- ❌ Database dumps with real data
- ❌ Any file containing real user data

**3. Generated/Build Files**

- ❌ `node_modules/` (can be reinstalled)
- ❌ `.next/`, `dist/`, `build/` (build artifacts)
- ❌ `*.log` files
- ❌ `coverage/` (test coverage reports)

**4. OS & IDE Files**

- ❌ `.DS_Store` (macOS)
- ❌ `Thumbs.db` (Windows)
- ❌ Personal VS Code settings (`.vscode/settings.json` with personal configs)

---

## ✅ Safe to Push

**1. Source Code**

- ✅ All `.ts`, `.tsx`, `.go`, `.jsx`, `.js` files
- ✅ React components, hooks, utilities
- ✅ Go handlers, services, repositories
- ✅ CSS, Tailwind configs

**2. Configuration Templates**

- ✅ `.env.example` (without real values)
- ✅ `package.json`, `go.mod`, `tsconfig.json`
- ✅ Docker configs (`Dockerfile`, `compose.yaml`)
- ✅ CI/CD workflows (`.github/workflows/`)

**3. Documentation**

- ✅ `README.md`
- ✅ `/docs` folder
- ✅ Code comments

**4. Git Files**

- ✅ `.gitignore`
- ✅ `.gitattributes`

---

## 📋 Checklist for Client Projects

### Before First Push

- [ ] Review `.gitignore` is comprehensive
- [ ] Create `.env.example` without sensitive data
- [ ] Remove any hardcoded credentials from code
- [ ] Check no database files are tracked
- [ ] Scan for accidentally committed secrets

### Every Commit

- [ ] Run `git status` before committing
- [ ] Double-check changed files with `git diff`
- [ ] Ensure no `.env` files in staging area
- [ ] Review commit for sensitive data

---

## 🛡️ Best Practices for Client Security

### 1. Use Environment Variables

**Bad** ❌:

```typescript
const apiUrl = "https://api.clientname.com";
const apiKey = "sk_live_1234567890abcdef";
```

**Good** ✅:

```typescript
const apiUrl = process.env.NEXT_PUBLIC_API_URL;
const apiKey = process.env.API_KEY; // Never commit the actual .env
```

### 2. Create Template Files

Create `.env.example`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username_here
DB_PASSWORD=your_password_here
DB_NAME=database_name

# JWT Secret (generate a random string)
JWT_SECRET=change_this_to_secure_random_string

# API URLs
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### 3. Document Secrets Required

In README.md or `/docs/guides/DEVELOPMENT_SETUP.md`:

```markdown
## Required Environment Variables

Before running the project, copy `.env.example` to `.env` and fill in:

- `DB_PASSWORD`: Your PostgreSQL password
- `JWT_SECRET`: Generate with `openssl rand -base64 32`
```

### 4. Use Git Secrets Scanner (Optional but Recommended)

```bash
# Install git-secrets
brew install git-secrets  # macOS
# or
sudo apt install git-secrets  # Linux

# Setup in your repo
git secrets --install
git secrets --register-aws  # Scan for AWS keys
git secrets --add 'password:|secret:|api_key:'
```

---

## 🔍 Check for Leaked Secrets

### Scan Committed History

```bash
# Check if .env was ever committed
git log --all --full-history -- "**/.env"

# Find potential secrets in history
git log -p | grep -i 'password\|secret\|api_key'
```

### Use GitHub Secret Scanning

If using GitHub:

1. Go to repository Settings
2. Security & analysis
3. Enable "Secret scanning"

---

## 🚨 If You Accidentally Committed Secrets

### Option 1: Recent Commit (Not Pushed)

```bash
# Remove file from last commit
git reset HEAD~1
git add .  # Re-add files except secrets
git commit -m "Your commit message"
```

### Option 2: Already Pushed (Nuclear Option)

**WARNING**: This rewrites Git history!

```bash
# Remove file from entire history
git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch be/.env" \
  --prune-empty --tag-name-filter cat -- --all

# Force push (requires team coordination!)
git push origin --force --all
```

**After removing secrets**:

1. ⚠️ Rotate ALL exposed credentials immediately
2. Change passwords in production
3. Regenerate API keys
4. Update JWT secrets

---

## 📊 Repository Settings Recommendations

### For Private Client Repos

**GitHub Settings**:

- ✅ Make repository **Private**
- ✅ Enable "Require pull request reviews"
- ✅ Enable "Secret scanning"
- ✅ Disable "Allow merge commits" (use squash/rebase)
- ✅ Add `.gitignore` as protected file

### For Public Repos (Open Source)

If you decide to make it public later:

- ⚠️ **NEVER** include real client data
- ⚠️ Use completely fake/demo data
- ⚠️ Extra paranoid about credentials
- ⚠️ Consider separate "demo" branch

---

## 🎯 Current Project Status

Your `.gitignore` is already protecting:
✅ `be/.env`
✅ `fe/.env`
✅ `be/koskosan.db`
✅ `node_modules/`
✅ `.next/`

**Recommendations**:

1. ✅ Your `.gitignore` sudah bagus!
2. 📝 Buat `.env.example` untuk dokumentasi
3. 🔒 Pastikan tidak ada hardcoded credentials di code
4. 📋 Document required environment variables

---

## ✅ Final Checklist Before Sharing with Client

- [ ] All sensitive data is in `.env` files (not committed)
- [ ] `.gitignore` is comprehensive
- [ ] No hardcoded passwords/API keys in code
- [ ] Documentation is complete and professional
- [ ] Code is clean and commented
- [ ] No debug logs with sensitive info
- [ ] Database schema doesn't contain real user data
- [ ] README has clear setup instructions

---

## 📞 Emergency Contacts

If you suspect a security breach:

1. Rotate all credentials immediately
2. Check GitHub security alerts
3. Review access logs
4. Notify affected parties if necessary

---

**Remember**: Keamanan data klien adalah prioritas #1! 🔒
