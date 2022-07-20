from setuptools import setup, find_packages

setup(
    name="cominnek",
    version="1.3.1",
    packages=find_packages(),
    url="https://github.com/Minnek-Digital-Studio/cominnek",
    license="MIT License",
    install_requires=["colorama==0.4.5"],
    author="Isaac Martinez",
    author_email="isaac@minnekdigital.com",
    description="Commits and pull requests",
    python_requires='>=3.9.1',
    entry_points={
        'console_scripts': ['cominnek = src.main:main'],
    }
)
