from setuptools import setup, find_packages

setup(
    name="cominnek",
    version="1.2.3",
    packages=find_packages(),
    url="",
    license="MIT License",
    install_requires=[],
    author="Isaac Martinez",
    author_email="isaac@minnekdigital.com",
    description="Commits and pull requests",
    python_requires='>=3.9.1',
    entry_points={
        'console_scripts': ['cominnek = src.main:main'],
    }
)
